import * as vm from "@pulumi/azure-quickstart-compute";

const vmach = new vm.Virtualmachine("page", {
    name : "vm1",
    adminUsername : "admin123",
    adminPassword : "unify!23$",
    location : "westus",
    vmSize : "Standard_D2s_v3"
});


