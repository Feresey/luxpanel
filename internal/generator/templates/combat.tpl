{{- $time := `(?P<LogTime>\d{2}:\d{2}:\d{2}\.\d{3})`}}
{{- $cmbt := printf "%s%s" $time `\s+CMBT\s+\|`}}
{{- $damage := printf "%s%s" $cmbt `\s+Damage\s+`}}
{{- $heal := printf "%s%s" $cmbt `\s+Heal\s+`}}
{{- $kill := printf "%s%s" $cmbt `\s+Killed\s+`}}
{{- $name := `[a-zA-Z0-9_/-]+`}}
{{- $id := `-?\d+`}}
{{- $float := `-?\d+\.\d+`}}


damage: >-
    (?s)^{{$damage}}
{{- wrap "Initiator" $name}}\|{{wrap "InitiatorID" $id -}}\s+->\s+{{"" -}}
({{wrap "Recipient" $name}}|{{wrap "ObjectName" $name}}\({{wrap "ObjectOwner" $name}}\))\|{{wrap "RecipientID" $id -}}\s+
{{- wrap "DamageTotal" $float}}\s+{{`\(` -}}
(h:{{- wrap "DamageHull" $float}})\s+{{"" -}}
(s:{{- wrap "DamageShield" $float}}){{`\)\s+` -}}
({{wrap "ActionSource" `\(?[a-zA-Z0-9_/-]+\)?`}})?\s+{{"" -}}
{{wrap "DamageModifiers" `[A-Z\|_]+` -}}
(\s+{{wrap "FriendlyFire" "<FriendlyFire>"}})?\s*$

heal: >-
    (?s)^{{$heal}}
{{- wrap "Initiator" $name}}\|{{wrap "InitiatorID" $id -}}\s+->\s+{{"" -}}
({{wrap "Recipient" $name}}|{{wrap "ObjectName" $name}}\({{wrap "ObjectOwner" $name}}\))\|{{wrap "RecipientID" $id -}}\s+
{{- wrap "Heal" $float}}\s+
{{- wrap "ActionSource" `\(?[a-zA-Z0-9_/-]+\)?`}}?\s*$

kill: >-
    (?s)^{{$kill -}}
(
{{- "" -}}
    ({{wrap "RecipientName" $name}}\s+{{wrap "RecipientShip" $name}})
    {{- "" -}}
    |
    {{- "" -}}
    ({{wrap "RecipientObject" $name}}|{{wrap "RecipientObjectName" $name}}\({{wrap "RecipientObjectOwner" $name}}\))
{{- "" -}}
)\|{{wrap "RecipientID" $id}};\s+killer\s+
{{- wrap "Initiator" $name}}\|{{wrap "InitiatorID" $id -}}
(\s+{{wrap "ActionSource" `\(?[a-zA-Z0-9_/-]+\)?`}})?{{"" -}}
(\s+{{wrap "FriendlyFire" "<FriendlyFire>"}})?\s*$
