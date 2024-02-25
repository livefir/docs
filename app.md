 <!-- start hero -->

The **Fir** toolkit is designed for Go developers with moderate html/css & js skills who want to progressively build reactive web apps without mastering complex web frameworks. It includes a Go library and an Alpine.js plugin.

<!-- end hero -->

<!-- start section1.0 -->

## 1. Start with html/template

A Fir web page begins as a standard server-side rendered page which reloads to show the new state on user interaction. Click the buttons below to see it in action.

<!-- end section1.0 -->

<!-- start section1.1 -->

When the html form is submitted, its is handled on the server in `onEvent` functions registered for inc and dec. The action/formaction attribute must be of the format `/?event=inc` where `inc` is the event name. Notice there is **no javascript** in the page.
Go the the directory where you have these files and run:

```bash
go run counter.go
```

Open your browser and go to [http://localhost:9867](http://localhost:9867) to see the counter in action.

<!-- end section1.1 -->

<!-- start section2.0 -->

## 2. Enhance with [alpinejs](https://alpinejs.dev/)

Later we use fir's alpinejs plugin to enhance the form submission and receive the re-rendered template as an event. The event is handled by the `$fir.replace()` helper function which updates the inner content of the div on which the event listener is declared. Click the buttons below to see reactivity in action. Open this page in two tabs to see the changes in one tab reflected in the other.

<!-- end section2.0 -->

<!-- start section2.1 -->

As you can notice, the count value is updated without a page reload. The server side code remains unchanged.
Fir's magic expression `@fir:event-name:event-state::template-name` piggybacks on [alpinejs event binding syntax](https://alpinejs.dev/directives/on#custom-events) to declare [html/templates](https://pkg.go.dev/html/template) to be re-rendered on the server.

<!-- end section2.1 -->

<!-- start section2.2 -->

```html
<div 
    @fir:inc:ok::count="$fir.replace()"
    @fir:dec:ok::count="$fir.replace()">
    {{ block "count" . }}
        <div>Count: {{ .count }}</div>
    {{ end }}
</div>
```

If the handler response status for event `inc` is `ok` then re-render the template named `count` on the server and return the html output to the event listener as a [CustomEvent](https://developer.mozilla.org/en-US/docs/Web/API/CustomEvent). The [CustomEvent.detail](https://developer.mozilla.org/en-US/docs/Web/API/CustomEvent/detail) property is used by the alpinejs plugin helper `$fir.replace()` to update the div on which the listener is declared.

<!-- end section2.2 -->

<!-- start section2.3 -->

Alternatively there is a short-hand form for wiring up multiple events with the same action.

```html
<div @fir:[inc:ok,dec:ok]::count="$fir.replace()">
    {{ block "count" . }}
    <div>Count: {{ .count }}</div>
    {{ end }}
</div>
```

Moreover, `fir` is able to automatically extract a template from elements on which fir events are declared. The above snippet can be further simplified.

```html
<div @fir:[inc:ok,dec:ok]="$fir.replace()">
   Count: {{ .count }}
</div>
```

Go the the directory where you have these files and run:

```bash
go run counter.go
```

Open your browser and go to [http://localhost:9867](http://localhost:9867) to see the reactive counter in action.

<!-- end section2.3 -->
