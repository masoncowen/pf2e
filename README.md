# TUI GM-aid for Pathfinder 2e
## Python version
Uses:
- Log how long encounters take
- Keep track of initiative
- Keep track of Character and Creature AC
- Quickly start pre-prepared combat encouters
- Track dice rolls for fun(?) statistics after a session or campaign (how many times did bless actually matter?) (todo)

## Go to-implement from Python version
- Utility Commands
    - [x] PASS (Not really a command, just used to pass commands to engines)
    - [ ] HELP
    - [ ] TODO
    - [x] QUIT (minimal version, just select Quit on start menu currently)
    - [ ] ECHO (not certain use-case, used in one-shot session for quick notes)
    - [x] BREAK (very minimal version, just records a timestamp and simple message. Not recorded to any log currently)
    - [ ] SAVE (wasn't even implemented in Python)
    - [ ] LOAD (wasn't even implemented in Python)
    - [ ] BEGIN
        - [ ] COMBAT (only half-implemented engine in Python)
- Combat Engine
    - [ ] Initiative tracker
    - [ ] Health management
    - [ ] Encounter system (it was hard coded in Python)
    - [ ] Party system (it was hard coded in Python)
## General todo
- [x] Load data from pathbuilder json files
