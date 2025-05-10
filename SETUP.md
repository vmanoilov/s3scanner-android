# S3Scanner Android App Setup Guide

This guide provides step-by-step instructions to complete the S3Scanner Android app build. The core scanning logic is implemented in Go, with an Android GUI frontend.

## Prerequisites

- Go 1.16 or later
- Android Studio 4.2 or later
- Android SDK Platform 30 or later
- Java Development Kit (JDK) 11
- Gomobile tool

## Quick Setup

### 1. Install Gomobile
```bash
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
```

### 2. Generate Android Bindings
```bash
# From project root
cd mobile/scanner
gomobile bind -target=android -o scanner.aar .
cp scanner.aar ../../app/libs/
```

### 3. Android Studio Setup
1. Open Android Studio
2. Click "Open Project"
3. Select the `app` directory
4. Wait for Gradle sync to complete

### 4. Configure Build
1. Open `app/build.gradle`
2. Verify the following dependencies:
```gradle
dependencies {
    implementation fileTree(dir: 'libs', include: ['*.aar'])
    implementation 'androidx.appcompat:appcompat:1.4.1'
    implementation 'com.google.android.material:material:1.5.0'
    implementation 'androidx.constraintlayout:constraintlayout:2.1.3'
}
```

### 5. Build and Run
1. Click "Build → Clean Project"
2. Click "Build → Rebuild Project"
3. Connect Android device or start emulator
4. Click "Run → Run 'app'"

## Troubleshooting

### Common Issues

1. **Binding Generation Failed**
   - Ensure GOPATH is set correctly
   - Run `go mod tidy` before binding
   - Check Go version compatibility

2. **Android Studio Sync Failed**
   - Update Gradle version
   - Invalidate caches/restart
   - Check SDK installation

3. **Runtime Crashes**
   - Verify .aar file is in app/libs
   - Check Java interface matches Go types
   - Enable debug logging

## Next Steps

1. Test the app thoroughly
2. Add custom providers if needed
3. Configure ProGuard rules
4. Prepare for release

## Support

For technical issues:
1. Check error logs
2. Search existing issues
3. Create detailed bug report

## Quick Reference

Key files:
```
app/libs/scanner.aar         # Generated Go bindings
app/build.gradle            # Build configuration
AndroidManifest.xml         # App permissions
MainActivity.java           # Main UI logic
S3Scanner.java             # Go code interface
```

Important commands:
```bash
# Regenerate bindings
gomobile bind -target=android -o scanner.aar .

# Clean build
./gradlew clean
./gradlew build

# Run tests
./gradlew test