# latin-replacement

Telegram bot: detects Uzbek Cyrillic messages, transliterates to Latin, deletes original, sends replacement.

## Structure

```
main.go                  # Bot init + message loop
transliterate/
  transliterate.go       # Cyrillic→Latin conversion library
go.mod                   # Module: github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
```

## Message Flow

1. Receive update → skip non-messages
2. `transliterate.HasCyrillic()` → skip if no Cyrillic
3. `transliterate.Do()` → convert text
4. Delete original message (needs admin "Delete Messages" perm)
5. Send replacement: `"{user} aytmoqchi bo'ldiki: {latin}\n{user} lotincha yoz olipta !"`
6. If original was reply → preserve `ReplyToMessageID`

## Key Functions

### main.go
- `main()` — bot init, update loop
- `senderName(u *tgbotapi.User) string` — returns `@username` or `FirstName LastName` or `"Someone"`

### transliterate/transliterate.go
- `HasCyrillic(s string) bool` — detects any Cyrillic rune, early exit
- `Do(text string) string` — digraph replace (НГ→NG) then rune-by-rune via `table`
- `table map[rune]string` — full Cyrillic→Latin map incl. Uzbek-specific (Ў,Қ,Ғ,Ҳ)
- `multiReplace []struct{from,to string}` — НГ/Нг/нг digraph rules

## Config

**Bot token hardcoded** in `main.go` — security issue, should move to env var `BOT_TOKEN`.  
No env vars, no config files currently used.

## Build & Run

```bash
go build -o latin-replacement .
./latin-replacement
```

Bot logs: `Authorized as @{bot_username}` on start.

## Known Limitations / TODOs

- Token hardcoded (should use `os.Getenv("BOT_TOKEN")`)
- No tests
- No graceful shutdown
- No rate limiting
- Deletion silently fails if bot lacks admin perms (logs error, continues)
- Go version: 1.25.6
