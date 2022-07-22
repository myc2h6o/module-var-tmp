# module-var-tmp

## Usage
```
go build
.\module-var-tmp.exe -path .\test
```

## Example output
```
[ERROR] test\main.tf:Line:20 Column:11; Property: tag; Variable: tag_a
[ERROR] test\main.tf:Line:21 Column:11; Property: tag; Variable: tag_b
[WARNING] (Ignored): test\main.tf:Line:9 Column:29; Property: name; Variable: name_prefix
[INFO] test\main.tf:Line:9 Column:48; Property: name; Variable: aks_name
[INFO] test\main.tf:Line:12 Column:26; Property: principal_id; Variable: role_assignment_principal_id
[INFO] test\main.tf:Line:14 Column:33; Property: conditional_prop; Variable: my_conditional_prop
[ERROR] test\main.tf:Line:14 Column:59; Property: conditional_prop; Variable: my_conditional_prop_2
[ERROR] test\main.tf:Line:16 Column:21; Property: nested_var_prop; Variable: props
[INFO] test\main.tf:Line:17 Column:16; Property: array_prop; Variable: my_array_prop
[INFO] test\main.tf:Line:25 Column:11; Property: n_a; Variable: n_a
[INFO] test\main.tf:Line:27 Column:8; Property: n_c; Variable: nested_prop_n_c
[ERROR] test\main.tf:Line:28 Column:8; Property: n_d; Variable: nested_prop_n_d_invalid
[INFO] test\sub-module\compute.tf:Line:7 Column:10; Property: name; Variable: resource_group_name
[ERROR] test\sub-module\compute.tf:Line:16 Column:19; Property: keepers; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:29 Column:30; Property: tags; Variable: tags
[INFO] test\sub-module\compute.tf:Line:40 Column:38; Property: delete_os_disk_on_termination; Variable: delete_os_disk_on_termination
[INFO] test\sub-module\compute.tf:Line:41 Column:38; Property: delete_data_disks_on_termination; Variable: delete_data_disks_on_termination
[INFO] test\sub-module\compute.tf:Line:134 Column:10; Property: tags; Variable: tags
[INFO] test\sub-module\compute.tf:Line:34 Column:41; Property: name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:38 Column:38; Property: vm_size; Variable: vm_size
[INFO] test\sub-module\compute.tf:Line:46 Column:14; Property: type; Variable: identity_type
[INFO] test\sub-module\compute.tf:Line:53 Column:22; Property: type; Variable: identity_type
[INFO] test\sub-module\compute.tf:Line:54 Column:53; Property: identity_ids; Variable: identity_ids
[INFO] test\sub-module\compute.tf:Line:59 Column:17; Property: id; Variable: vm_os_id
[INFO] test\sub-module\compute.tf:Line:60 Column:47; Property: publisher; Variable: vm_os_publisher
[INFO] test\sub-module\compute.tf:Line:61 Column:47; Property: offer; Variable: vm_os_offer
[INFO] test\sub-module\compute.tf:Line:62 Column:47; Property: sku; Variable: vm_os_sku
[INFO] test\sub-module\compute.tf:Line:63 Column:38; Property: version; Variable: vm_os_version
[INFO] test\sub-module\compute.tf:Line:67 Column:35; Property: name; Variable: vm_hostname
[ERROR] test\sub-module\compute.tf:Line:70 Column:25; Property: managed_disk_type; Variable: storage_account_type
[INFO] test\sub-module\compute.tf:Line:79 Column:27; Property: disk_size_gb; Variable: data_disk_size_gb
[ERROR] test\sub-module\compute.tf:Line:80 Column:27; Property: managed_disk_type; Variable: data_sa_type
[INFO] test\sub-module\compute.tf:Line:76 Column:30; Property: name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:87 Column:30; Property: name; Variable: vm_hostname
[WARNING] (Ignored): test\sub-module\compute.tf:Line:89 Column:51; Property: lun; Variable: nb_data_disk
[ERROR] test\sub-module\compute.tf:Line:91 Column:27; Property: managed_disk_type; Variable: data_sa_type
[ERROR] test\sub-module\compute.tf:Line:96 Column:25; Property: computer_name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:97 Column:22; Property: admin_username; Variable: admin_username
[INFO] test\sub-module\compute.tf:Line:98 Column:22; Property: admin_password; Variable: admin_password
[INFO] test\sub-module\compute.tf:Line:99 Column:22; Property: custom_data; Variable: custom_data
[ERROR] test\sub-module\compute.tf:Line:103 Column:39; Property: disable_password_authentication; Variable: enable_ssh_key
[ERROR] test\sub-module\compute.tf:Line:108 Column:29; Property: path; Variable: admin_username
[ERROR] test\sub-module\compute.tf:Line:116 Column:29; Property: path; Variable: admin_username
[ERROR] test\sub-module\compute.tf:Line:137 Column:19; Property: enabled; Variable: boot_diagnostics
[INFO] test\sub-module\compute.tf:Line:144 Column:38; Property: name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:148 Column:35; Property: vm_size; Variable: vm_size
[INFO] test\sub-module\compute.tf:Line:150 Column:35; Property: delete_os_disk_on_termination; Variable: delete_os_disk_on_termination
[INFO] test\sub-module\compute.tf:Line:151 Column:35; Property: license_type; Variable: license_type
[INFO] test\sub-module\compute.tf:Line:211 Column:10; Property: tags; Variable: tags
[INFO] test\sub-module\compute.tf:Line:156 Column:14; Property: type; Variable: identity_type
[INFO] test\sub-module\compute.tf:Line:163 Column:22; Property: type; Variable: identity_type
[INFO] test\sub-module\compute.tf:Line:164 Column:53; Property: identity_ids; Variable: identity_ids
[INFO] test\sub-module\compute.tf:Line:169 Column:17; Property: id; Variable: vm_os_id
[INFO] test\sub-module\compute.tf:Line:170 Column:47; Property: publisher; Variable: vm_os_publisher
[INFO] test\sub-module\compute.tf:Line:171 Column:47; Property: offer; Variable: vm_os_offer
[INFO] test\sub-module\compute.tf:Line:172 Column:47; Property: sku; Variable: vm_os_sku
[INFO] test\sub-module\compute.tf:Line:173 Column:38; Property: version; Variable: vm_os_version
[ERROR] test\sub-module\compute.tf:Line:180 Column:25; Property: managed_disk_type; Variable: storage_account_type
[INFO] test\sub-module\compute.tf:Line:177 Column:28; Property: name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:189 Column:27; Property: disk_size_gb; Variable: data_disk_size_gb
[ERROR] test\sub-module\compute.tf:Line:190 Column:27; Property: managed_disk_type; Variable: data_sa_type
[INFO] test\sub-module\compute.tf:Line:186 Column:30; Property: name; Variable: vm_hostname
[ERROR] test\sub-module\compute.tf:Line:201 Column:27; Property: managed_disk_type; Variable: data_sa_type
[INFO] test\sub-module\compute.tf:Line:197 Column:30; Property: name; Variable: vm_hostname
[WARNING] (Ignored): test\sub-module\compute.tf:Line:199 Column:51; Property: lun; Variable: nb_data_disk
[ERROR] test\sub-module\compute.tf:Line:206 Column:25; Property: computer_name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:207 Column:22; Property: admin_username; Variable: admin_username
[INFO] test\sub-module\compute.tf:Line:208 Column:22; Property: admin_password; Variable: admin_password
[ERROR] test\sub-module\compute.tf:Line:230 Column:19; Property: enabled; Variable: boot_diagnostics
[INFO] test\sub-module\compute.tf:Line:242 Column:34; Property: tags; Variable: tags
[INFO] test\sub-module\compute.tf:Line:236 Column:37; Property: name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:247 Column:28; Property: name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:250 Column:25; Property: allocation_method; Variable: allocation_method
[INFO] test\sub-module\compute.tf:Line:251 Column:25; Property: sku; Variable: public_ip_sku
[INFO] test\sub-module\compute.tf:Line:253 Column:25; Property: tags; Variable: tags
[INFO] test\sub-module\compute.tf:Line:265 Column:28; Property: name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:269 Column:10; Property: tags; Variable: tags
[ERROR] test\sub-module\compute.tf:Line:274 Column:58; Property: name; Variable: remote_port
[INFO] test\sub-module\compute.tf:Line:283 Column:33; Property: source_address_prefixes; Variable: source_address_prefixes
[INFO] test\sub-module\compute.tf:Line:293 Column:35; Property: enable_accelerated_networking; Variable: enable_accelerated_networking
[INFO] test\sub-module\compute.tf:Line:302 Column:10; Property: tags; Variable: tags
[INFO] test\sub-module\compute.tf:Line:290 Column:38; Property: name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:296 Column:40; Property: name; Variable: vm_hostname
[INFO] test\sub-module\compute.tf:Line:297 Column:37; Property: subnet_id; Variable: vnet_subnet_id
```
