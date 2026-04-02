## Difficulty Retargeting

Bitcoin adjusts its difficulty every 2,016 blocks, roughly every two weeks. The goal is to maintain an average block interval of 10 minutes.

At each retargeting point, nodes calculate how long the previous 2,016 blocks took. If blocks came faster than 10 minutes on average, the target decreases (difficulty rises). If blocks came slower, the target increases (difficulty drops). The formula compares actual elapsed time to the expected 20,160 minutes.

A safeguard prevents the difficulty from changing by more than a factor of four in a single retarget. This limits how fast difficulty can spike or plummet. The mechanism is entirely algorithmic. No committee votes on it. Every node computes the same new target from the same chain data.
