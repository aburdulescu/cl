# cl
`cl` is a tool that takes one or more group(s) of color/pattern(regular expression)
and applies these rules on the given input(a file or piped through stdin) by
coloring all lines(or characters, depending on the selected mode(see [Modes](https://gitlab.com/aburdulescu/cl/blob/master/README.md#modes)))
that match the given pattern with the given color.

## Modes
### Line mode(default mode)
In line mode, `cl` will color the whole line that matches the given pattern with
the given color.

### Match mode
In line mode, `cl` will color only the characters that match the given pattern
with the given color.