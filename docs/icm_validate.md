## icm validate

Validate intermodal container markings

### Synopsis

Validate intermodal container markings.

Configuration for separators is generated first time you
execute a command that requires the configuration.

Flags for output formatting can overridden with a config file.
Edit default configuration:

  $HOME/.icm/config.yml

```
icm validate [flags]
```

### Examples

```
icm validate ABC
icm validate ABC --pattern container-number
icm validate ABC U
icm validate --sep-owner-equip '' --sep-serial-check '-' ABC U 123456 0
icm validate ABC U 123456 0 20G1
icm validate 20G1
icm generate | icm validate
icm generate --count 10 | icm validate
icm generate --count 10 | icm validate --output fancy
```

### Options

```
  -p, --pattern string            sets pattern matching to
                                                      auto = matches automatically a pattern
                                          container-number = matches a container number
                                                     owner = matches a three letter owner code
                                  owner-equipment-category = matches a three letter owner code with equipment category ID
                                                 size-type = matches length, width+height and type code
                                   (default "auto")
      --output string             sets output to
                                   auto = for a single line 'fancy' and for multiple lines 'csv' output 
                                    csv = machine readable CSV output
                                  fancy = human readable fancy output
                                   (default "auto")
      --no-header                 omits header of CSV output
      --sep-owner-equip string    ABC(x)U1234560   20G1  (x) separates owner code and equipment category id (default " ")
      --sep-equip-serial string   ABCU(x)1234560   20G1  (x) separates equipment category id and serial number (default " ")
      --sep-serial-check string   ABCU123456(x)0   20G1  (x) separates serial number and check digit (default " ")
      --sep-check-size string     ABCU1234560 (x)  20G1  (x) separates check digit and size (default "   ")
      --sep-size-type string      ABCU1234560   20(x)G1  (x) separates size and type (default " ")
  -h, --help                      help for validate
```

### SEE ALSO

* [icm](icm.md)	 - Validate or generate intermodal container markings
