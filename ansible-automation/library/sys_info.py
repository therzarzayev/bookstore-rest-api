#!/usr/bin/python

from ansible.module_utils.basic import AnsibleModule
import platform

def main():
    module = AnsibleModule(argument_spec={})
    system_info = {
        'platform': platform.system(),
        'platform_release': platform.release(),
        'platform_version': platform.version(),
    }
    module.exit_json(changed=False, system_info=system_info)

if __name__ == '__main__':
    main()
