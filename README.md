# exa CLI

A Go command-line interface for the [Exa AI](https://exa.ai) search API — neural web search built for AI applications.

## Installation

```bash
git clone <repo>
cd exa-cli
go build -o ~/bin/exa .
```

Or run directly without installing:

```bash
go run . <command> [flags]
```

## Authentication

Set your API key as an environment variable:

```bash
export EXA_API_KEY=your-api-key-here
```

Or pass it inline with every command:

```bash
exa --api-key <key> search "your query"
```

Get an API key at [dashboard.exa.ai](https://dashboard.exa.ai).

---

## Commands

### `search` — Web search

Search the web using Exa's neural or keyword search engine.

```bash
exa search "large language models 2025"
```

| Flag | Default | Description |
|------|---------|-------------|
| `-n, --num-results` | `10` | Number of results |
| `--type` | `auto` | Search mode: `auto`, `neural`, `fast`, `deep` |
| `--category` | — | Content type: `news`, `research paper`, `company`, `pdf`, `tweet`, `personal site`, `financial report`, `people` |
| `--start-date` | — | Published after (YYYY-MM-DD) |
| `--end-date` | — | Published before (YYYY-MM-DD) |
| `--include-domains` | — | Comma-separated domains to include |
| `--exclude-domains` | — | Comma-separated domains to exclude |
| `--include-text` | — | Results must contain this text |
| `--exclude-text` | — | Results must not contain this text |
| `--text` | `false` | Include full page text |
| `--summary` | `false` | Include AI-generated summary |
| `--highlights` | `false` | Include highlighted excerpts |
| `--livecrawl` | — | `always`, `fallback`, `never`, `auto` |
| `--max-age` | — | Max content age in hours |
| `--json` | `false` | Output raw JSON |

**Examples:**

```bash
# Recent AI news with summaries
exa search "OpenAI news" --category news --start-date 2025-01-01 --summary

# Deep neural search on academic sites
exa search "transformer attention mechanisms" \
  --type neural \
  --include-domains arxiv.org,scholar.google.com \
  --highlights

# Pipe to jq
exa search "Go 1.24 features" --json | jq '.[].url'
```

---

### `find-similar` — Find similar pages

Find web pages similar to a given URL.

```bash
exa find-similar https://example.com/article
```

| Flag | Default | Description |
|------|---------|-------------|
| `-n, --num-results` | `10` | Number of results |
| `--exclude-source-domain` | `false` | Exclude the input URL's domain from results |
| `--include-domains` | — | Restrict to these domains |
| `--exclude-domains` | — | Exclude these domains |
| `--start-date` | — | Published after (YYYY-MM-DD) |
| `--end-date` | — | Published before (YYYY-MM-DD) |
| `--text` | `false` | Include full page text |
| `--summary` | `false` | Include AI-generated summary |
| `--highlights` | `false` | Include highlighted excerpts |
| `--livecrawl` | — | Live crawl mode |
| `--max-age` | — | Max content age in hours |
| `--json` | `false` | Output raw JSON |

**Examples:**

```bash
# Find similar articles, excluding the source domain
exa find-similar https://techcrunch.com/2025/01/article --exclude-source-domain

# Find similar research papers
exa find-similar https://arxiv.org/abs/2401.12345 --summary -n 5
```

---

### `get-contents` — Retrieve page content

Retrieve the full text, summary, or highlights of one or more URLs.

```bash
exa get-contents https://example.com
exa get-contents https://site1.com https://site2.com
```

| Flag | Default | Description |
|------|---------|-------------|
| `--text` | `true` | Include page text |
| `--summary` | `false` | Include AI-generated summary |
| `--highlights` | `false` | Include highlighted excerpts |
| `--livecrawl` | — | Live crawl mode |
| `--max-age` | — | Max content age in hours |
| `--json` | `false` | Output raw JSON |

**Examples:**

```bash
# Get text of a documentation page
exa get-contents https://pkg.go.dev/net/http

# Always live-crawl for freshest content
exa get-contents https://example.com --livecrawl always

# Get summaries for multiple URLs
exa get-contents https://site1.com https://site2.com --summary
```

---

### `answer` — AI-generated answer with citations

Ask a question and get a direct answer backed by live web sources.

```bash
exa answer "What is the Exa API?"
```

| Flag | Default | Description |
|------|---------|-------------|
| `-m, --model` | `exa` | Model: `exa` or `exa-pro` (more thorough) |
| `--system-prompt` | — | Custom system prompt |
| `--location` | — | User location for localised results |
| `--text` | `false` | Include full source text |
| `--json` | `false` | Output raw JSON (answer + citations array) |

**Examples:**

```bash
# Quick factual question
exa answer "Who won the Nobel Prize in Physics 2024?"

# Detailed answer with exa-pro
exa answer "Compare RAG vs fine-tuning for LLMs" --model exa-pro

# Localised result
exa answer "Best coffee shops near me" --location "Paris, France"

# JSON for programmatic use
exa answer "Latest Go release features" --json | jq '.answer'
```

---

### `research` — Deep research task

Run a multi-step research task and receive a structured, comprehensive result.

```bash
exa research "Competitive landscape of AI search engines in 2025"
```

| Flag | Default | Description |
|------|---------|-------------|
| `--json` | `false` | Output raw JSON |

**Examples:**

```bash
# Research a topic
exa research "State of quantum computing startups in 2025"

# Research with JSON output
exa research "Top Go web frameworks with pros and cons" --json
```

---

## Global flags

These flags apply to every command:

| Flag | Description |
|------|-------------|
| `--api-key <key>` | Exa API key (or set `EXA_API_KEY`) |
| `--base-url <url>` | Override API base URL |

---

## Output

By default all commands print human-readable, formatted output. Add `--json` to any command to get raw JSON — useful for piping to `jq` or other tools.

```bash
exa search "golang concurrency" --json | jq '.[0]'
```

---

## Project structure

```
exa-cli/
├── main.go            # Entry point
├── go.mod / go.sum
├── exa/
│   ├── client.go      # HTTP client
│   └── types.go       # Request / response types
└── cmd/
    ├── root.go        # Root command, auth wiring, shared helpers
    ├── context.go     # Cobra context helpers
    ├── output.go      # Pretty-printer
    ├── search.go      # exa search
    ├── similar.go     # exa find-similar
    ├── contents.go    # exa get-contents
    ├── answer.go      # exa answer
    └── research.go    # exa research
```

---

## Based on

[exa-py](https://github.com/exa-labs/exa-py) — the official Exa Python SDK.
