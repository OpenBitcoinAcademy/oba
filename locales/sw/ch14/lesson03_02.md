## Kujenga Mti wa Hati

Mti wa Merkle unajengwa kutoka hash za TapLeaf na TapBranch.

Kila jani: tagged_hash("TapLeaf", leaf_version || compact_size(script_length) || script). Toleo la jani kwa Tapscript ya sasa ni 0xC0.

Kila tawi: tagged_hash("TapBranch", sorted(left_hash, right_hash)). Watoto wawili wanapangwa kwa leksikografia kabla ya kuhesabu hash, kuhakikisha mti una muundo mmoja wa kawaida bila kujali mpangilio wa kuingiza.

Mti unaweza kuwa na viwango hadi 128 vya kina, kuruhusu majani hadi 2^128. Miti ya vitendo ni ndogo zaidi. Hati zilizowekwa kwenye kina kirefu (karibu na mzizi) zina uthibitisho mfupi wa Merkle na zinagharimu kidogo kutumia. Weka mbadala unaowezekana zaidi kwenye kina 1. Weka njia za dharura zinazotumika nadra zaidi kina zaidi.
