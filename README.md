# RSS Scraper

## Overview

The RSS Scraper is a simple Go project designed to fetch and parse RSS feeds from various websites. This tool allows users to easily aggregate and read RSS feeds in a structured format.

## Features

- Fetch RSS feeds from multiple sources.
- Parse and extract useful information such as title, link, publication date, and summary.
- Save the extracted information in a structured format (e.g., JSON or CSV).
- Command-line interface for ease of use.

## Requirements

- Go 1.0 +
- `requests`
- `feedparser`

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/rss-scraper.git
    cd rss-scraper
    ```

2. Install the required packages:
    ```bash
    go mod tidy
    ```

3. Run the Project:
    ```bash
    make run 
    ```

## Usage

To use the RSS Scraper, follow these steps:

1. Create a file named `feeds.txt` in the root directory of the project. This file should contain the URLs of the RSS feeds you want to scrape, one per line.

    Example `feeds.txt`:
    ```
    https://example.com/feed
    https://anotherexample.com/rss
    ```

2. Run the scraper:
    ```bash
    Go scraper.py
    ```

3. The scraped data will be saved in a file named `output.json` (or `output.csv` depending on your preference).

## Configuration

You can configure the output format (JSON or CSV) by modifying the `config.py` file:

```Go
