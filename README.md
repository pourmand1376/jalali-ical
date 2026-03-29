Persian calendar .ics file 

- 2025
  https://ramin.tech/jalali-ical/2025_(1403-1404)_persian_calendar.ics
- 2026
  https://ramin.tech/jalali-ical/2026_(1404-1405)_persian_calendar.ics

### How to create the iCalendar file

```bash
go run cmd/jalali-ical.go --year 2026
```

-> output: `./2026_(1404-1405)_persian_calendar.ics`

### Garbage collection calendars

Running the command also generates two additional calendars for garbage collection reminders:

- **Odd-day**: events on days where the Jalali day-of-month is odd (1, 3, 5, …, 29)
- **Even-day**: events on days where the Jalali day-of-month is even (2, 4, 6, …, 30)

Each event is at **21:00 Tehran time** with a **10-minute** duration. Day 31 is excluded from both calendars because not all Jalali months have 31 days, making it an inconsistent collection day.
