# RSS Aggregator API

This project is a fully functional RSS aggregator built using Go and Chi Router.

## Table of Contents

- [RSS Aggregator API](#rss-aggregator-api)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Installation](#installation)
    - [Prerequisites](#prerequisites)
    - [Steps](#steps)
  - [Usage](#usage)
    - [Running the Application](#running-the-application)
    - [API Endpoints](#api-endpoints)
  - [Configuration](#configuration)
  - [Contributing](#contributing)
  - [License](#license)

## Introduction

The RSS Aggregator is designed to collect and organize posts from multiple RSS feeds. It periodically fetches new posts and provides a simple API to interact with the aggregated data.

## Features

- **Fetch and store RSS feeds:** Automatically fetches RSS feeds and stores them in a database.
- **User management:** Users can follow different feeds and see posts from those feeds.
- **API:** A simple and intuitive API to interact with the aggregated posts and feeds.
- **Concurrency:** Efficiently handles multiple feeds concurrently using goroutines.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)

### Steps

1. **Clone the repository:**

```bash
git clone <your-repo-url>
cd rss-aggregator
```

2. **Install dependencies:**

```bash
go mod tidy
```

3. **Set up the database:**

Ensure PostgreSQL is running and create a database for the project. Update the database configuration in the `config` file or environment variables.

4. **Run database migrations:**

```bash
goose up
```

5. **Build and run the application:**

```bash
go build
./rss-aggregator
```

## Usage

### Running the Application

To start the application, use the following command:

```bash
./rss-aggregator
```

### API Endpoints

- **GET /feeds:** Get a list of all available feeds.
- **POST /feeds:** Create a new feed.
- **GET /feeds/:id:** Get details of a specific feed.
- **POST /feed-follows:** Follow a feed.
- **GET /feed-follows:** Get a list of feeds followed by the user.
- **GET /posts:** Get a list of posts from followed feeds.

## Configuration

Configuration options can be set through environment variables or a configuration file. The following are key configuration options:

- `DATABASE_URL`: URL for the PostgreSQL database.
- `PORT`: Port on which the server will run.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
