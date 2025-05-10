# S3Scanner Android App

## Overview
S3Scanner is an Android application that helps security professionals and developers check for misconfigured cloud storage buckets. It provides a simple, user-friendly interface to scan buckets across multiple cloud providers and identify potential security risks.

## Features
- Scan cloud storage buckets from multiple providers:
  - Amazon S3
  - Google Cloud Storage
  - Azure Blob Storage
  - DigitalOcean Spaces
- Check for:
  - Bucket existence
  - Public accessibility
  - Region information
- User-friendly Material Design interface
- Quick, non-intrusive scanning
- Clear security warnings
- Support for custom providers

## Getting the APK

### Automatic Build with GitHub Actions
1. Fork this repository
2. Go to your forked repository on GitHub
3. Enable GitHub Actions if not already enabled
4. Go to the "Actions" tab
5. Select the "Android CI/CD" workflow
6. Click "Run workflow"
7. Once the workflow completes:
   - Click on the completed workflow run
   - Scroll down to "Artifacts"
   - Download either:
     - `app-debug.apk` for testing
     - `app-release.apk` for release version

### Manual Build
See [SETUP.md](SETUP.md) for detailed build instructions.

## Installation
1. Download the APK from GitHub Actions artifacts
2. Enable "Install from Unknown Sources" in Android settings
3. Install the APK
4. Grant required permissions when prompted

## Usage
1. Launch the app
2. Enter the bucket name to scan
3. Select the cloud provider
4. Tap "Scan Bucket"
5. Review the results

## Security Note
This tool is intended for security auditing purposes only. Always ensure you have proper authorization before scanning any buckets. The tool performs minimal, read-only operations to check bucket accessibility.

## Building from Source
Check [SETUP.md](SETUP.md) for detailed manual build instructions.

## Contributing
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments
- Built with Gomobile
- Uses Material Design components
- Original S3Scanner Go tool authors

## Troubleshooting GitHub Actions Build
If the build fails:
1. Check the workflow run logs for specific errors
2. Verify Go version compatibility
3. Ensure Android SDK components are available
4. Check if Gomobile bindings generated successfully
5. Look for any Gradle build issues
