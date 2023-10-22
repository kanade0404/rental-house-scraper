# Rental House Scraper

## System Design

### URL Frontier

URLをURL Filterから待ち受けて、検索すべきURLをHTML Downloaderに渡す。
本来はQueueから受けるが、初回はDatabaseから取得して実行する。

### HTML Downloader

URL Frontierから受け取ったURLを元にHTMLをダウンロードする。

### Contents Parser

HTML Downloaderから受け取ったHTMLをパースして、必要な情報を抽出する。

### URL Extractor

Contents Parserから受け取った情報から、次に検索すべきURLを抽出する。

### URL Filter

URLが検索すべきURLかどうかを判定する。
