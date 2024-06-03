# Data Clean Room TEE Base Image
The base image is used to build user's custom image that is going to run in google cloud confidential space.  This folder contains the tools that required by the base image including an encryption tool to encrypt the output `ipynb` file and an attestation tool to generate custom attestation report. 

## Build
Just run the `build.sh` to build the base image and push to gcp docker repository.