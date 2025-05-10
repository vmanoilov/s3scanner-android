package com.example.s3scanner;

import scanner.Scanner;
import scanner.MobileScanner;
import scanner.ScanResult;

/**
 * Java interface for the S3Scanner Go implementation
 */
public class S3Scanner {
    private final MobileScanner scanner;

    public static class ScanResult {
        private final String bucketName;
        private final String region;
        private final boolean exists;
        private final boolean isPublic;
        private final String error;

        public ScanResult(scanner.ScanResult goResult) {
            this.bucketName = goResult.getBucketName();
            this.region = goResult.getRegion();
            this.exists = goResult.getExists();
            this.isPublic = goResult.getIsPublic();
            this.error = goResult.getError() != null ? goResult.getError().getMessage() : null;
        }

        public String getBucketName() { return bucketName; }
        public String getRegion() { return region; }
        public boolean getExists() { return exists; }
        public boolean getIsPublic() { return isPublic; }
        public String getError() { return error; }
    }

    public S3Scanner() {
        this.scanner = Scanner.newMobileScanner();
    }

    public String[] getSupportedProviders() {
        return Scanner.getSupportedProviders();
    }

    public ScanResult scanBucket(String bucketName, String provider) {
        scanner.ScanResult goResult = scanner.scanBucket(bucketName, provider);
        return new ScanResult(goResult);
    }

    public String getVersion() {
        return scanner.getVersion();
    }
}