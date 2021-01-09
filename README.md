# tailwify
Simple auto-script tool for Auto-generating starter templates with TailwindCSS.

## Install
```
go get -u github.com/TheBoringDude/tailwify
```

## What it does?
Nothing fancy, just the following, .. 
- Install the specified framework
- Add TailwindCss packages and others
- Configure the framework (modify files / strings)

## Supported and Added Frameworks
- Next.JS   (framework: `next`)
- Gatsby    (framework: `gatsby`)
- Vue3 w/ Vite  (framework: `viteVue3`)

## How to Use:
```bash
tailwify generate {framework} -p {project-name}
```

#### Installations and configurations are based from the official docs: https://tailwindcss.com/docs/installation

### Image
![app screenshot](./screenshot.png)

#### &copy; TheBoringDude