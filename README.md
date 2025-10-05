# AWZ-Buddy 🚀

**AWZ-Buddy** is a lightweight CLI tool to help you **manage and audit your AWS resources**.  
Its main goal is to identify **unused or misconfigured resources**, prevent **cost leaks**, and keep your AWS environment **clean and secure**.

---

## ⚡ Features (initial version)

- **S3 Utilities**:
  - List all buckets
  - Detect publicly accessible buckets
  - List all empty buckets

- **Future plans**:
  - EC2:
    - Estimate instance costs based on type and usage
    - Check if EC2 instances are running but underutilized
  - Elastic IPs: detect unused addresses
  - CloudWatch metrics and cost reports
  - Export results to JSON/CSV for easy audits

---

## 🛠 Installation

AWZ-Buddy uses **Go** and can be run locally with **Nix**. Follow these steps:

### 1. Install [Nix](https://nixos.org/download/)
If you don't have Nix installed yet, run:

```bash
curl -L https://nixos.org/nix/install | sh
```

> After installation, restart your shell or source your profile


### 2. Clone the repository
```bash
git clone https://github.com/EfrainPerez23/awz-buddy.git
cd awz-buddy
```

## 3. Run nix develop
```bash
nix develop
go mod tidy
```

## 4. Run awz-buddy
```bash
make # this comand will create a binary called awz-buddy at dist folder
# Execute awz-buddy depending on your OS
./dist/awz-buddy-OS-ARCH
# For example:
./dist/awz-buddy-linux-amd64
./dist/awz-buddy-windows-amd64.exe
./dist/awz-buddy-darwin-amd64
```