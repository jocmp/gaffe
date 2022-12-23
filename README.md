# Gaffe

Gaffe presents an HTML list for the CSV output of [sivel/speedtest-cli](https://github.com/sivel/speedtest-cli).

I run this little web server on my NAS to keep track of a cron job that logs new speed test results over time using speedtest-cli.
The goal is to have a simple view that doesn't require manually copying a CSV and pasting it into Google Sheets.

![Index Example](./screenshots/index-example.md)

## Getting Started

To try it out locally, clone this repository

```sh
git clone https://github.com/jocmp/gaffe.git
```

Follow the instructions on [sivel/speedtest-cli](https://github.com/sivel/speedtest-cli) to create a speed test result.

```sh
python speedtest-cli >> speedtest.csv
```

Finally, run the debug build

```sh
export WEB_HOST=8700
export CSV_PATH=speedtest.csv
make run
```

You should now see the webpage on <http://localhost:8700>.


