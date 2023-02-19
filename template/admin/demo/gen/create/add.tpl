{{define "__PATH__"}}
  {{template "layout_form" .}}
{{end}}

{{define "content"}}
<form class="form-horizontal m" id="form-add">
___REPLACE___
</form>
{{end}}

{{define "script"}}
    <script>
        function submitHandler() {
            if ($.validate.form()) {
                $.operate.save(window.b5go.oasUrl, $('#form-add').serialize());
            }
        }
    </script>
{{end}}
