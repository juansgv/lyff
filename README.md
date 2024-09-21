# Project Q

## Overview

**Project Q** is a cloud-hosted call center application optimized for small businesses. Built using Rust, Actix-web, and integrated with Moshi Rust, the application leverages advanced voice processing and AI agents to deliver efficient call handling. Additionally, it features a 2D visualization inspired by Pokémon Emerald to represent agent interactions, enhancing user engagement and monitoring.

## Features

- **Voice Processing:** Transcribe and analyze incoming audio calls using integrated voice processing modules.
- **AI Agents:** Intelligent agents handle queries and provide responses based on transcriptions.
- **Cloud Deployment:** Hosted on Google Cloud Platform (GCP) for scalability and reliability.
- **Cost Optimization:** Configured to minimize operational costs, ideal for small businesses.
- **2D Visualization:** Interactive Pokémon Emerald-style graphics to visualize agent activities and system status.

## Technology Stack

- **Programming Language:** Rust
- **Web Framework:** Actix-web
- **AI Integration:** Moshi Rust
- **Cloud Hosting:** Google Cloud Platform (GCP)
- **Visualization:** Bevy (Rust game engine) or SDL2 for 2D graphics
- **Version Control:** GitHub

## Repository Structure

- `src/`: Contains the Rust source code for the application.
- `static/`: Static assets like images and audio files.
- `templates/`: HTML templates for rendering views.
- `Cargo.toml`: Rust package configuration.
- `Dockerfile`: Docker configuration for containerization.
- `docker-compose.yml`: Docker Compose configuration for running the application in a containerized environment.

Table of Contents
Prerequisites
Install Pulumi
Set Up Pulumi Project
Define Infrastructure
1. Choose AWS Regions
2. Set Up VPCs and Networking
3. Containerize Your Applications
4. Set Up ECR (Elastic Container Registry)
5. Deploy Kubernetes Clusters with EKS
6. Deploy Backend Services
7. Deploy Web UI
8. Configure Load Balancing and DNS
9. Set Up Monitoring and Logging
Deploy with Pulumi
Continuous Integration and Deployment (CI/CD)
Optimize for Latency and Reliability
Maintain and Scale
Complete Pulumi Example
Additional Resources
---
Prerequisites
Before we begin, ensure you have the following:
AWS Account: Sign up for an AWS account.
AWS CLI: Install and configure the AWS CLI.
Pulumi Account: Sign up for a Pulumi account.
Pulumi CLI: Install the Pulumi CLI.
Docker: Install Docker.
Node.js: Install Node.js (used for Pulumi with TypeScript/JavaScript).