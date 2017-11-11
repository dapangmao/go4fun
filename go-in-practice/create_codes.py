import os

def make_index():
    strs = """
    1. Getting in to Go
    1.1. What is Go?
    1.2. Noteworthy Aspects of Go
    1.2.1. Multiple Return Values
    1.2.2. A Modern Standard Library
    1.2.3. Concurrency with goroutines and channels
    1.2.4. Go the toolchain, more than a language
    1.3. Go in the vast language landscape
    1.3.1. C and Go
    1.3.2. Go and Java
    1.3.3. Python, PHP, and Go
    1.3.4. JavaScript, node.js, and Go
    1.4. Getting up and running in Go
    1.4.1. Installing Go
    1.4.2. Git, Mercurial, and Version Control
    1.4.3. Workspace
    1.4.4. Environment Variables
    1.5. Hello Go
    2. A Solid Foundation
    2.1. CLI Applications, the Go Way
    2.1.1. Command Line Flags
    2.1.2. Command Line Frameworks
    2.2. Handling Configuration
    2.3. Working with Real World Web Servers
    2.3.1. Starting up and shutting down a server
    2.3.2. Routing Web Requests
    3. Concurrency in Go
    3.1. Understanding Go&#39;s Concurrency Model
    3.2. Working with Goroutines
    3.3. Working with Channels
    4. Handling Errors and Panics
    4.1. Error Handling
    4.2. The Panic System
    4.2.1. Differentiating panics from errors
    4.2.2. Working with Panics
    4.2.3. Recovering from Panics
    4.2.4. Panics and Goroutines
    5. Debugging and Testing
    5.1. Locating Bugs
    5.1.1. Wait, Where Is My Debugger?
    5.2. Logging
    5.2.1. Using Go&#8217;s Logger
    5.2.2. Working with System Loggers
    5.3. Accessing Stack Traces
    5.4. Testing
    5.4.1. Unit Testing
    5.4.2. Generative Testing
    5.5. Performance Tests and Benchmarks
    6. HTML and Email Template Patterns
    6.1. Working with HTML templates
    6.1.1. Standard library HTML package overview
    6.1.2. Adding Functionality Inside Templates
    6.1.3. Limiting template parsing
    6.1.4. When template execution breaks
    6.1.5. Mixing Templates
    6.2. Using Templates for Email
    7. Serving and Receiving, Assets and Forms
    7.1. Serving static content
    7.2. Handling Form Posts
    7.2.1. Introduction to form requests
    7.2.2. Working with files and miltipart submissions
    7.2.3. Working with raw multipart data
    8. Working with Web Services
    8.1. Using REST APIs
    8.1.1. Using The HTTP Client
    8.1.2. When Faults Happen
    8.2. Passing And Handling Errors Over HTTP
    8.2.1. Generating Custom Errors
    8.2.2. Reading And Using Custom Errors
    8.3. Parsing And Mapping JSON
    8.4. Versioning REST APIs
    9. Using the Cloud
    9.1. What Is Cloud Computing?
    9.1.1. The Different Types of Cloud
    9.1.2. Containers and Cloud Native Applications
    9.2. Managing Cloud Services
    9.2.1. Avoid Cloud Provider Lock&#151;in
    9.2.2. Dealing with Divergent Errors
    9.3. Running On Cloud Servers
    9.3.1. Runtime Detection
    9.3.2. Building For The Cloud
    9.3.3. Runtime Monitoring
    10. Communication Between Cloud Services
    10.1. Microservices and High Availability
    10.2. Communicating Between Services
    10.2.1. Making REST Faster
    10.2.2. Moving Beyond REST
    11. Reflection and Code Generation
    11.1. Three features of reflection
    11.2. Structs, tags, and annotations
    11.2.1. Annotating structs
    11.2.2. Using tag annotations
    11.3. Generating Go code with Go code
    """
    index = {}
    prev = None
    for l in strs.split('\n'):
        l = l.strip()
        if not l:
            continue
        key, value = l.split('. ')
        if '.' not in key:
            index[key] = '## Chapter {} '.format(key) + value
            prev = key
        else:
            index[prev] += '\n' + '- ' + key + ' ' + value
    return index


def get_init_codes():
    res = {}

    for root, dirs, files in os.walk(".", topdown=False):
        if '.git' in root or '.ipynb' in root:
            continue
        for name in files:
            if name.endswith('.go'):
                current = ""
                with open(os.path.join(root, name), 'r') as infile:
                    for i, l in enumerate(infile):
                        current += l
                res[root + '/' + name ] = current 
    return res

def get_codes_by_chapter(init):
    res = {}
    prev = ""
    for k, v in init.items():
        ks = k.split('/')[1:]
        key = ks[0].replace("chapter", "")
        value = "#### " + " -> ".join(ks) + "\n```go\n" + v + "```\n"
        try:
            _ = int(key)
            if key not in res:
                res[key] = value
            else:
                res[key] += value
        except:
            print(key)
            pass
    return res


def output_md(codes, index):
    s = set()

    with open('./codes.md', 'w') as out:
        for k in sorted(codes, key=lambda x: int(x)):
            if k not in s:
                print("\n{}\n".format(index[k]), file=out)
                s.add(k)
            print(codes[k], file=out)    
    
if __name__ == "__main__":

    index = make_index()
    init = get_init_codes()
    cbc = get_codes_by_chapter(init)
    output_md(cbc, index)
            

