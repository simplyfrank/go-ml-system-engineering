Command-line applications are among the easiest ways to handle user input and output. This chapter will focus on command-line-based interactions, such as command-line arguments, configuration, and environment variables. We will conclude with a library for coloring text output in Unix and Bash for Windows.

With the recipes in this chapter, you should be equipped to handle expected and unexpected user input. The Catching and handling signals recipe is an example of cases where users may send unexpected signals to your application, and the pipes recipe is a good alternative to taking user inputs compared to flags or command-line arguments.

The ANSI color recipe will hopefully provide some examples of cleaning up output to users. For example, in logging, being able to color text based on its purpose can sometimes make large blocks of text significantly clearer.

In this chapter, we will cover the following recipes:

- Using command-line flags
- Using command-line arguments
- Reading and setting environment variables
- Configuration using TOML, YAML, and JSON
- Working with Unix pipes