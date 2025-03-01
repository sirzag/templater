# Templater

Templater is a Go-based CLI tool for generating files from templates. It provides a simple way to create reusable templates with configurable options and generates files based on user input.

## Features

- Create and manage file templates in a centralized location
- Define template variables through interactive prompts
- Support for various input types (string, number, boolean, enum)
- Generate files with proper extension handling
- Automatic directory creation

## Installation

```bash
go install github.com/sirzag/templater/cmd/templater@latest
```

## Usage

Once installed, you can use the templater command followed by the template name:

```bash
templater <template-command>
```

For example:

```bash
templater component
```

This will prompt you for the necessary information to generate a file based on the "component" template.

## Creating Templates

Templates are stored in the `~/.templater/templates` directory. Each template consists of:

1. A template file (e.g. `component.tmpl`)
2. A configuration file (e.g. `component.config.json` or `component.config.jsonc`)

### Configuration File Structure

```jsonc
{
  "cmd": "component",
  "description": "Create a React component",
  "out": "{WORKING_DIR}/",
  "template": "./component.tmpl", // relative to ~/.templater/templates. I know, not obvious yet
  "extension": ".jsx",
  "options": [
    {
      "name": "COMPONENT_TYPE",
      "type": "enum",
      "description": "Component type",
      "required": true,
      "default": "Functional",
      "values": ["Functional", "Class"]
    },
    {
      "name": "WITH_STYLES",
      "type": "boolean",
      "description": "Include styles?",
      "default": true
    }
  ]
}
```

### Template File Example

```jsx
import React from 'react';
{{if .WITH_STYLES}}
import styles from './{{.FILE_NAME}}.module.css';
{{end}}

{{if eq .COMPONENT_TYPE "Functional"}}
const {{.FILE_NAME}} = (props) => {
  return (
    <div>
      {{.FILE_NAME}} Component
    </div>
  );
};
{{else}}
class {{.FILE_NAME}} extends React.Component {
  render() {
    return (
      <div>
        {{.FILE_NAME}} Component
      </div>
    );
  }
}
{{end}}

export default {{.FILE_NAME}};
```

## Template Options

Templater supports the following option types:

- `string` - Text input
- `number` - Numeric input
- `boolean` - Yes/No selection
- `enum` - Selection from a list of predefined values

Each option can have:
- `name` - Variable name to use in the template
- `description` - Prompt text
- `required` - Whether input is required
- `default` - Default value
- `values` - For enum types, list of available options

## Built-in Variables

- `FILE_NAME` - The name of the file being created (without extension)

## Development

To build the project from source:

```bash
git clone https://github.com/sirzag/templater.git
cd templater
go build -o templater ./cmd/templater
```

## License

MIT
