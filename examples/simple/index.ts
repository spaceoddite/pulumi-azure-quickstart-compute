import * as xyz from "@pulumi/azure-quickstart-compute";

const page = new xyz.StaticPage("page", {
    indexContent: "<html><body><p>Hello world!</p></body></html>",
});

export const bucket = page.bucket;
export const url = page.websiteUrl;
