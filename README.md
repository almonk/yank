# @yank

### Introduction

@yank lets you treat standard CSS classes like @mixins in Sass.

For example, given `tachyons.css` as our **definitions** file, we could write the following @yank rules in a standard css file `myclasses.css`;

```
.my-button {
  @yank .bg-blue;
  @yank .white;
  @yank .link;
  @yank .pa2;
  @yank .br2;
}

.my-input {
  @yank .bg-white;
  @yank .ba;
  @yank .b--gray;
  @yank .dark-gray;
  @yank .br2;
  @yank .pa2;
}
```

[You can view this example on jsfiddle](https://jsfiddle.net/almonk/dvdzjvLa/).

Running `./yank` on the command line will compile this file into the following `myclasses--compiled.css`:

```
/* This file was compiled with @yank */
/* See the docs at https://github.com/almonk/yank */
.my-button {
  background-color:#357edd;
  color:#fff;
  text-decoration:none;
  padding:.5rem;
  border-radius:.25rem;
}
.my-input {
  background-color:#fff;
  border-style:solid;
  border-width:1px;
  border-color:#777;
  color:#333;
  border-radius:.25rem;
  padding:.5rem;
}
```

### Usage
You can run these on the command line as;
`./yank -input=example/myclasses.css -definitions=example/mycss.css -output=example/myclasses--compiled.css`

---
### Help

Running `./yank --help` will show you the flags you can pass to @yank

```
  -definitions string
    	File that @yank uses to expand your classes (default "mycss.css")
  -input string
    	File where your @yank rules are defined (default "myclasses.css")
  -output string
    	CSS file with compiled @yank rules (default "myclasses--compiled.css")
```
