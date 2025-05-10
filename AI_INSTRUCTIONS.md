# AI Instructions for S3Scanner Android App Completion

## Project Overview
This is a conversion of a Go-based S3 bucket scanner into an Android GUI application. The project uses Gomobile to bind Go code to Android and provides a user-friendly interface for scanning cloud storage buckets.

## Project Structure

```
/mobile/scanner/
  ├── scanner.go     # Core scanning logic and types
  └── mobile.go      # Mobile-specific implementations

/app/src/main/
  ├── java/com/example/s3scanner/
  │   ├── MainActivity.java    # Main Android UI
  │   └── S3Scanner.java      # Java interface to Go code
  └── res/layout/
      └── activity_main.xml    # UI layout
```

## Components to Implement

1. Go-Android Integration
   - Generate Gomobile bindings from the scanner package
   - Place generated .aar file in app/libs/
   - Update build.gradle to include the native library

2. Android Project Setup
   - Configure Gradle dependencies
   - Set up proper permissions in AndroidManifest.xml
   - Implement error handling for native code

3. Testing & Validation
   - Implement UI tests
   - Add input validation
   - Test native code integration
   - Verify proper error handling

## Key Integration Points

1. Scanner Package (Go):
   ```go
   package scanner
   
   type MobileScanner struct {...}
   func NewMobileScanner() *MobileScanner
   func (s *MobileScanner) ScanBucket(bucketName, provider string) *ScanResult
   ```

2. Java Interface:
   ```java
   public class S3Scanner {
       private final MobileScanner scanner;
       public ScanResult scanBucket(String bucketName, String provider)
   }
   ```

## Required Dependencies

In build.gradle:
```gradle
dependencies {
    implementation fileTree(dir: 'libs', include: ['*.aar'])
    implementation 'com.google.android.material:material:1.5.0'
    implementation 'androidx.constraintlayout:constraintlayout:2.1.3'
}
```

## Build Instructions

1. Generate Gomobile bindings:
   ```bash
   go get golang.org/x/mobile/cmd/gomobile
   gomobile init
   gomobile bind -target=android -o scanner.aar ./mobile/scanner
   ```

2. Android build:
   ```bash
   ./gradlew assembleDebug
   ```

## Testing Strategy

1. Unit Tests
   - Test Go scanner functionality
   - Test Java interface methods
   - Validate input handling

2. Integration Tests
   - Verify Go-Java communication
   - Test error propagation
   - Check memory management

3. UI Tests
   - Validate user input handling
   - Test progress indicators
   - Verify error messages

## Completion Checklist

1. [ ] Generate and verify Gomobile bindings
2. [ ] Set up Android project with all dependencies
3. [ ] Implement remaining error handling
4. [ ] Add input validation
5. [ ] Set up CI/CD pipeline
6. [ ] Complete testing suite
7. [ ] Generate release build