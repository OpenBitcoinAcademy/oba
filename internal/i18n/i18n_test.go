package i18n

import (
	"sort"
	"testing"
	"testing/fstest"
)

func testFS() fstest.MapFS {
	return fstest.MapFS{
		"en/ui.toml": &fstest.MapFile{
			Data: []byte(`
[nav]
home = "Home"
settings = "Settings"
next = "Next"
progress = "{{completed}} of {{total}} complete"

[quiz]
remaining_one = "{{count}} question remaining"
remaining_other = "{{count}} questions remaining"
`),
		},
		"es/ui.toml": &fstest.MapFile{
			Data: []byte(`
[nav]
home = "Inicio"
settings = "Configuracion"
next = "Siguiente"
progress = "{{completed}} de {{total}} completo"

[quiz]
remaining_one = "{{count}} pregunta restante"
remaining_other = "{{count}} preguntas restantes"
`),
		},
	}
}

func setup(t *testing.T) {
	t.Helper()
	if err := Init(testFS()); err != nil {
		t.Fatalf("Init: %v", err)
	}
	SetLocale("en")
}

func TestT_EnglishKey(t *testing.T) {
	setup(t)
	got := T("nav.home")
	if got != "Home" {
		t.Errorf("T(nav.home) = %q, want %q", got, "Home")
	}
}

func TestT_MissingKeyReturnsKey(t *testing.T) {
	setup(t)
	got := T("nonexistent.key")
	if got != "nonexistent.key" {
		t.Errorf("T(nonexistent.key) = %q, want the key itself", got)
	}
}

func TestT_SpanishLocale(t *testing.T) {
	setup(t)
	SetLocale("es")
	got := T("nav.home")
	if got != "Inicio" {
		t.Errorf("T(nav.home) with es locale = %q, want %q", got, "Inicio")
	}
}

func TestT_FallbackToEnglish(t *testing.T) {
	setup(t)
	// Spanish locale doesn't have every key, but in our test data it does.
	// Add a key only in English to test fallback.
	catalogs["en"]["nav.special"] = "Special"
	SetLocale("es")
	got := T("nav.special")
	if got != "Special" {
		t.Errorf("T(nav.special) with es locale = %q, want English fallback %q", got, "Special")
	}
}

func TestTFmt_Interpolation(t *testing.T) {
	setup(t)
	got := TFmt("nav.progress", map[string]string{
		"completed": "3",
		"total":     "10",
	})
	want := "3 of 10 complete"
	if got != want {
		t.Errorf("TFmt(nav.progress) = %q, want %q", got, want)
	}
}

func TestTFmt_SpanishInterpolation(t *testing.T) {
	setup(t)
	SetLocale("es")
	got := TFmt("nav.progress", map[string]string{
		"completed": "3",
		"total":     "10",
	})
	want := "3 de 10 completo"
	if got != want {
		t.Errorf("TFmt(nav.progress) es = %q, want %q", got, want)
	}
}

func TestTPlural_One(t *testing.T) {
	setup(t)
	got := TPlural("quiz.remaining", 1, map[string]string{"count": "1"})
	want := "1 question remaining"
	if got != want {
		t.Errorf("TPlural(quiz.remaining, 1) = %q, want %q", got, want)
	}
}

func TestTPlural_Other(t *testing.T) {
	setup(t)
	got := TPlural("quiz.remaining", 5, map[string]string{"count": "5"})
	want := "5 questions remaining"
	if got != want {
		t.Errorf("TPlural(quiz.remaining, 5) = %q, want %q", got, want)
	}
}

func TestTPlural_Zero(t *testing.T) {
	setup(t)
	// Zero uses _other form in English.
	got := TPlural("quiz.remaining", 0, map[string]string{"count": "0"})
	want := "0 questions remaining"
	if got != want {
		t.Errorf("TPlural(quiz.remaining, 0) = %q, want %q", got, want)
	}
}

func TestTPlural_SpanishOne(t *testing.T) {
	setup(t)
	SetLocale("es")
	got := TPlural("quiz.remaining", 1, map[string]string{"count": "1"})
	want := "1 pregunta restante"
	if got != want {
		t.Errorf("TPlural(quiz.remaining, 1) es = %q, want %q", got, want)
	}
}

func TestSetLocale_UnknownFallsBackToEnglish(t *testing.T) {
	setup(t)
	SetLocale("xx")
	got := Locale()
	if got != "en" {
		t.Errorf("Locale() after SetLocale(xx) = %q, want %q", got, "en")
	}
	// Verify T still works with English.
	val := T("nav.home")
	if val != "Home" {
		t.Errorf("T(nav.home) after unknown locale = %q, want %q", val, "Home")
	}
}

func TestAvailable(t *testing.T) {
	setup(t)
	got := Available()
	sort.Strings(got)
	if len(got) != 2 || got[0] != "en" || got[1] != "es" {
		t.Errorf("Available() = %v, want [en es]", got)
	}
}

func TestInit_MissingEnglish(t *testing.T) {
	noEnglish := fstest.MapFS{
		"es/ui.toml": &fstest.MapFile{
			Data: []byte(`[nav]
home = "Inicio"
`),
		},
	}
	err := Init(noEnglish)
	if err == nil {
		t.Error("Init without en/ui.toml should return error")
	}
}

func TestInit_MalformedTOML(t *testing.T) {
	bad := fstest.MapFS{
		"en/ui.toml": &fstest.MapFile{
			Data: []byte(`this is not valid toml {{{}}`),
		},
	}
	err := Init(bad)
	if err == nil {
		t.Error("Init with malformed TOML should return error")
	}
}

func TestTPlural_FallbackToEnglishPlural(t *testing.T) {
	setup(t)
	// Add a plural key only in English, not in Spanish.
	catalogs["en"]["items_one"] = "{{count}} item"
	catalogs["en"]["items_other"] = "{{count}} items"
	SetLocale("es")

	got := TPlural("items", 1, map[string]string{"count": "1"})
	if got != "1 item" {
		t.Errorf("TPlural fallback to en _one = %q, want %q", got, "1 item")
	}

	got = TPlural("items", 5, map[string]string{"count": "5"})
	if got != "5 items" {
		t.Errorf("TPlural fallback to en _other = %q, want %q", got, "5 items")
	}
}

func TestTPlural_MissingOneFormFallsToOther(t *testing.T) {
	setup(t)
	// A locale that only has _other, not _one.
	catalogs["en"]["things_other"] = "{{count}} things"
	// No things_one defined.

	got := TPlural("things", 1, map[string]string{"count": "1"})
	if got != "1 things" {
		t.Errorf("TPlural missing _one should fall to _other = %q, want %q", got, "1 things")
	}
}

func TestTPlural_MissingKeyReturnsKey(t *testing.T) {
	setup(t)
	got := TPlural("nonexistent", 3, map[string]string{"count": "3"})
	if got != "nonexistent" {
		t.Errorf("TPlural missing key = %q, want key itself", got)
	}
}

func TestTPlural_SpanishFallbackOtherForm(t *testing.T) {
	setup(t)
	// Spanish has _other but not _one for this key.
	catalogs["es"]["widgets_other"] = "{{count}} widgets (es)"
	SetLocale("es")

	got := TPlural("widgets", 1, map[string]string{"count": "1"})
	if got != "1 widgets (es)" {
		t.Errorf("TPlural es missing _one falls to es _other = %q, want %q", got, "1 widgets (es)")
	}
}

func TestInterpolate_NoArgs(t *testing.T) {
	got := interpolate("Hello world", nil)
	if got != "Hello world" {
		t.Errorf("interpolate with nil args = %q, want %q", got, "Hello world")
	}
}

func TestInterpolate_MissingPlaceholder(t *testing.T) {
	// Placeholder not in args stays as-is.
	got := interpolate("Hello {{name}}", map[string]string{})
	if got != "Hello {{name}}" {
		t.Errorf("interpolate with missing arg = %q, want placeholder preserved", got)
	}
}
