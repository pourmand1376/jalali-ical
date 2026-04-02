Persian calendar .ics file 

- 2025
  https://ramin.tech/jalali-ical/2025_(1403-1404)_persian_calendar.ics
- 2026
  https://ramin.tech/jalali-ical/2026_(1404-1405)_persian_calendar.ics

### How to use in Google Calendar

1. **Create a new Google Calendar**: Go to [Google Calendar](https://calendar.google.com), click the **+** next to "Other calendars" in the left sidebar, and select **Create new calendar**. Give it a name (e.g. "Persian Calendar") and click **Create calendar**.

2. **Import the .ics file**: In Google Calendar settings, select your newly created calendar, scroll down to **Import & export**, choose the `.ics` file you downloaded, and click **Import**.

3. **Set event notifications**: Open Google Calendar settings for the imported calendar, go to **Event notifications**, and add two reminders:
   - **60 minutes** (1 hour) before
   - **30 minutes** before

### How to create the iCalendar file

```bash
go run cmd/jalali-ical.go --year 2026
```

-> output: `./2026_(1404-1405)_persian_calendar.ics`

### Garbage collection calendars

Running the command also generates two additional calendars for garbage collection reminders:

- **Odd-day**: events on days where the Jalali day-of-month is odd (1, 3, 5, …, 29)
- **Even-day**: events on days where the Jalali day-of-month is even (2, 4, 6, …, 30)

Each event is at **22:00 Tehran time** with a **30-minute** duration. Day 31 is excluded from both calendars because garbage is not collected on that day in our town.

### Releasing calendars

To publish calendars for the next 5 years, create and push a tag (any name):

```bash
git tag v1.0.0
git push origin v1.0.0
```

The **Release Calendars** GitHub Actions workflow will automatically:
1. Generate `.ics` files for the current year and the following year.
2. Upload all generated calendar files as assets to the corresponding GitHub Release.
