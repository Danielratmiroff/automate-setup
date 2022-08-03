## Automate-Setup

Simple project to install and configure all the things the way I like in my Ubuntu machine.

## Install

You can simply clone the repo into your local machine and run the `automate-setup` file in your terminal

```
./automate-setup
```

I recommend this way if you don't have go installed already and you fully trust me. (_might be risky or maybe not?_ 🙄)

#### Or, you can go the manual way:

You'll need to [Install Go](https://go.dev/doc/install) first

```bash
https://github.com/Danielratmiroff/automate-setup.git
cd automate-setup
go install # (or just go build -o .)
```

**Please note**: _If you get an error claiming that automate-setup cannot be found or is not defined, you may need to add ~/go/bin to your $PATH (MacOS/Linux), or %HOME%\go\bin (Windows)._

## Usage

> As of now, it asumes your linux username is "daniel". _(will add support for dynamic usernames in the near future)_

Run it to install and configure everything at once.

Ideal if you have a new laptop or just happends that you just deleted everything by accident _(I'm that dumb)_

```bash
./automate-setup
```

You can install individual pkgs or software by passing parameters.

```bash
./automate-setup exa # Will install only the "exa" package.
```

## What will be installed?

**Software**

- ansible
- lazygit
- neovim
- zshrc
- docker

**Dotfiles**

- neovim
- fish
- oh-my-fish
- zshrc
- oh-my-zshrc

**Packages**

- git
- unzip
- gzip
- curl
- wget
- fd-find
- ripgrep
- pandoc
- pipenv
- progress
- findutils
- gawk
- npm
- nodejs
- fish
- oh-my-fish
- peco
- tmux
- rsync
- python3.8-venv
- bat
- direnv
- exa
- dust
- yarn
- golang

_Any better software alternative suggestions are very much welcome_
