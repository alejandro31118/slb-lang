### Why SLB?
- `S`: Simple
- `L`: Lightweight
- `B`: Bad designed

### Why Go?
The main reason why I chose Go for this is beacuse is a very simple language with features that will satisfies the needs of this project.
I thougth about C too, but why bother myself with the memory management if this is just a challenge.

### Parts
- `Token`: Lang reserved keywords. Each token holds information about its type and the actual content it represents.
- `Lexer`: It takes the source code as input and breaks it down into individual tokens. It scans through the code, identifies patterns, and assigns appropriate token types to them.
- `AST (Abstract syntax tree)`: It is a structured representation of the code's grammar and logic. It helps the computer understand the relationships between different parts of the code, like which operations should happen first and how expressions are connected.
- `Parser`: It takes the tokens from the lexer and arranges them into the abstract syntax tree. It ensures that the code follows the correct grammar and rules of the language, turning the raw tokens into a structured tree that can be understood and processed.
- `REPL (Read–Eval–Print-Loop)`: It's a tool that lets you interact with your language in real-time. You type in a command or code snippet, the REPL reads it, evaluates it (runs the code), and then prints out the result. This loop allows you to experiment, test, and explore the language without having to write and run full programs.
