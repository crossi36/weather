weather is a CLI tool for [chubin/wttr.in](https://github.com/chubin/wttr.in).

Installation is a easy as

    go get github.com/crossi36/weather

## Usage

    $ weather berlin
    Weather report: Berlin, Germany

        \   /     Sunny
         .-.      12-14 °C
      ― (   ) ―   ← 19 km/h
         `-’      10 km
        /   \     0.0 mm

wttr.in offers a variety of translations, which can also be used with this tool.
You can set the language you prefer with the environment variable `weather_lang`.
The default language is English.

    $ env weather_lang=de weather berlin
    Wetterbericht für: Berlin, Germany

        \   /     Sonnig
         .-.      12-14 °C
      ― (   ) ―   ← 19 km/h
         `-’      10 km
        /   \     0.0 mm

You can also search for places by prefixing your query with a tilde:

    $ weather ~eiffel tower
    Weather report: eiffel tower

        \  /       Partly cloudy
      _ /"".-.     17 °C
        \_(   ).   ↑ 7 km/h
        /(___(__)  10 km
                   0.0 mm

You can also use IP-addresses (direct) or domain names (prefixed with @)
as a location specificator:

    $ weather @github.com
    Weather report: San Francisco, United States of America

        \   /     Clear
         .-.      18 °C
      ― (   ) ―   → 11 km/h
         `-’      16 km
        /   \     0.0 mm

    $ weather 8.8.8.8
    Weather report: Mountain View, United States of America

        \   /     Clear
         .-.      20 °C
      ― (   ) ―   ↓ 0 km/h
         `-’      16 km
        /   \     0.0 mm

You can use 3-letters airport codes if you want to get the weather information
about some airports:

    $ weather muc
    Weather report: muc, Munich International Airport, Germany

        \  /       Partly cloudy
      _ /"".-.     13 °C
        \_(   ).   ↑ 5 km/h
        /(___(__)  20 km
                   0.0 mm

weather can also be used to get the current moon phase

    $ weather moon
                        -------.	 
                            . .   `--.	 
                            .       . `-.	 
                            .  @@@@@      `-.	 
                            @@@@@@@   .    \	 
                            @@@@@@@       . \.	 
                        .   @@@@@@@   O      \	 
                            @@@@@@@@@@     @@@ \	 
                        . @@@@@@@@@@@@@ o @@@@|	 
                            @@@@@@@@@@@@    @@  \	 First Quarter +
                        o     @@@@@@@@ @@@@    |	 0  7:03:35
                            .  @@   . @@@@@@@  |	 Full Moon -
                        .-.     @@@   @@@@@@@  |	 7  8:43:11
                        `-'   . @@@@   @@@@  o /	 
                                    @@   .       |	 
                            .  O    .     o   . /	 
                                .    .       /	 
                        _   .      .   .-. /'	 
                                        `-' /	 
                            o    O   .   .-'	 
                                .    .-'	 
                            .      .--'	 
                        -------'
or for a specific date

    $ weather moon@2017-09-06
                    .------------.	 
                .--'  o     . .   `--.	 
            .-'   .    O   .       . `-.	 
            -'@   @@@@@@@   .  @@@@@      `-.	 
        /@@@  @@@@@@@@@@@   @@@@@@@   .    \	 
        ./    o @@@@@@@@@@@   @@@@@@@       . \.	 
    /@@  o   @@@@@@@@@@@.   @@@@@@@   O      \	 
    @@@@   .   @@@@@@@o    @@@@@@@@@@     @@@ \	 
    |@@@@@               . @@@@@@@@@@@@@ o @@@@|	 
    /@@@@@  O  `.-./  .      @@@@@@@@@@@@    @@  \	 First Quarter +
    | @@@@    --`-'       o     @@@@@@@@ @@@@    |	 7 12:45:40
    |@ @@@        `    o      .  @@   . @@@@@@@  |	 Full Moon -
    |       @@  @         .-.     @@@   @@@@@@@  |	 0 10:04:55
    \  . @        @@@     `-'   . @@@@   @@@@  o /	 
    |      @@   @@@@@ .           @@   .       |	 
            @@@@  @\@@    /  .  O    .     o   . /	 
    \  o  @@     \ \  /         .    .       /	 
        `\     .    .\.-.___   .      .   .-. /'	 
        \           `-'                `-' /	 
            -.   o   / |     o    O   .   .-'	 
            `-.   /     .       .    .-'	 
                `--.       .      .--'	 
                    `------------'	