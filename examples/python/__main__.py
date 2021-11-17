"""An Azure RM Python Pulumi program"""
# python -m pip install -e location\to\pulumi-azure-quickstart-compute\sdk\python


from pulumi_azure_quickstart_compute import virtualmachine as vm

# Create an Azure virtual machine

vm1 = vm.Virtualmachine("vm1",vm.VirtualmachineArgs(name="workload1",admin_username="admin123",admin_password="Unify123$",location="westus",vm_size="Standard_D2s_v3",image_type="windows"))
