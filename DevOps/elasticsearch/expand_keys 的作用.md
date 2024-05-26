https://www.elastic.co/guide/en/beats/filebeat/8.13/decode-json-fields.html

(Optional) A Boolean value that specifies whether keys in the decoded JSON should be recursively de-dotted and expanded into a hierarchical object structure. For example, `{"a.b.c": 123}` would be expanded into `{"a":{"b":{"c":123}}}`.

其实很清晰，实践一下也确实如此，但是文档中显示没什么出入