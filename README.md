## Goland

First pass at a Wayland library for Go.

Currently working on code generation using the Wayland protocol spec XML
document and message codec/dispatch.

```testing/wayland_pipe``` contains a tool that will connect to an existing
compositor (the socket must be named "weston") and provides a new display socket
called "compositor". It'll just copy the messages between clients and the real
display server and print their contents to the console. Right now, all it
actually interprets is the message header.
