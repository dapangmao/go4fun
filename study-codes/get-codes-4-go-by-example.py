import os
import re

index = """
Hello World
Values
Variables
Constants
For
If/Else
Switch
Arrays
Slices
Maps
Range
Functions
Multiple Return Values
Variadic Functions
Closures
Recursion
Pointers
Structs
Methods
Interfaces
Errors
Goroutines
Channels
Channel Buffering
Channel Synchronization
Channel Directions
Select
Timeouts
Non-Blocking Channel Operations
Closing Channels
Range over Channels
Timers
Tickers
Worker Pools
Rate Limiting
Atomic Counters
Mutexes
Stateful Goroutines
Sorting
Sorting by Functions
Panic
Defer
Collection Functions
String Functions
String Formatting
Regular Expressions
JSON
Time
Epoch
Time Formatting / Parsing
Random Numbers
Number Parsing
URL Parsing
SHA1 Hashes
Base64 Encoding
Reading Files
Writing Files
Line Filters
Command-Line Arguments
Command-Line Flags
Environment Variables
Spawning Processes
Exec'ing Processes
Signals
Exit
"""

res = {}
for root, dirs, files in os.walk("/Users/admin/projects/gobyexample/examples", topdown=False):
    if '.git' in root or '.ipynb' in root:
        continue
    for name in files:
        if name.endswith('.go'):
            current = ""
            with open(os.path.join(root, name), 'r') as infile:
                for i, l in enumerate(infile):
                    current += l
            res[name ] = current


index = index.replace("'", "")
idx = ['-'.join(re.findall(r"[\w']+", x.lower())) for x in index.split('\n') if x]

with open("go-by-example.md", "w") as outfile:
    for x in idx:
        print("- {}\n".format(x), file=outfile)
        current_codes = res[x+'.go']
        print('```go\n{}\n```\n'.format(current_codes), file=outfile)
