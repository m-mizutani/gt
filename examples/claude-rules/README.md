# Claude Code Rules for gt

This directory contains a [Claude Code rule](https://docs.anthropic.com/en/docs/claude-code/memory) template for projects using the gt test library.

## Usage

Copy `test-gt.md` to your project's `.claude/rules/` directory:

```bash
mkdir -p .claude/rules
cp test-gt.md /path/to/your-project/.claude/rules/
```

Or create a symlink to share across multiple projects:

```bash
mkdir -p .claude/rules
ln -s /path/to/gt/examples/claude-rules/test-gt.md .claude/rules/test-gt.md
```

## How it works

The rule file uses YAML frontmatter to apply only to test files:

```yaml
---
paths:
  - "**/*_test.go"
---
```

When Claude Code works with `*_test.go` files, it will automatically load these rules and use gt correctly.

## Customization

You can customize the rules for your project:

- Add project-specific test patterns
- Remove methods you don't use
- Add additional guidelines

## More information

- [gt documentation](https://pkg.go.dev/github.com/m-mizutani/gt)
- [Claude Code memory documentation](https://docs.anthropic.com/en/docs/claude-code/memory)
