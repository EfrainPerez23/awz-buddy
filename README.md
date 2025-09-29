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

```bash
git clone https://github.com/EfrainPerez23/awz-buddy.git
cd awz-buddy
go mod tidy
go build -o awz-buddy
```