package com.example.s3scanner;

import android.os.Bundle;
import android.view.View;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ProgressBar;
import android.widget.Spinner;
import android.widget.TextView;
import androidx.appcompat.app.AppCompatActivity;
import com.google.android.material.snackbar.Snackbar;

public class MainActivity extends AppCompatActivity {
    private S3Scanner scanner;
    private EditText bucketNameInput;
    private Spinner providerSpinner;
    private TextView resultText;
    private Button scanButton;
    private ProgressBar progressBar;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        // Initialize scanner
        scanner = new S3Scanner();

        // Initialize UI components
        bucketNameInput = findViewById(R.id.bucket_name_input);
        providerSpinner = findViewById(R.id.provider_spinner);
        resultText = findViewById(R.id.result_text);
        scanButton = findViewById(R.id.scan_button);
        progressBar = findViewById(R.id.progress_bar);

        // Set up provider spinner
        ArrayAdapter<String> providerAdapter = new ArrayAdapter<>(
            this,
            android.R.layout.simple_spinner_item,
            scanner.getSupportedProviders()
        );
        providerAdapter.setDropDownViewResource(android.R.layout.simple_spinner_dropdown_item);
        providerSpinner.setAdapter(providerAdapter);

        scanButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                String bucketName = bucketNameInput.getText().toString().trim();
                
                if (bucketName.isEmpty()) {
                    showError("Please enter a bucket name");
                    return;
                }

                // Disable button and show progress
                scanButton.setEnabled(false);
                progressBar.setVisibility(View.VISIBLE);
                resultText.setText("");

                String provider = providerSpinner.getSelectedItem().toString();

                new Thread(new Runnable() {
                    @Override
                    public void run() {
                        try {
                            final S3Scanner.ScanResult result = scanner.scanBucket(bucketName, provider);
                            runOnUiThread(new Runnable() {
                                @Override
                                public void run() {
                                    displayResult(result);
                                    scanButton.setEnabled(true);
                                    progressBar.setVisibility(View.GONE);
                                }
                            });
                        } catch (Exception e) {
                            runOnUiThread(new Runnable() {
                                @Override
                                public void run() {
                                    showError("Scan failed: " + e.getMessage());
                                    scanButton.setEnabled(true);
                                    progressBar.setVisibility(View.GONE);
                                }
                            });
                        }
                    }
                }).start();
            }
        });
    }

    private void displayResult(S3Scanner.ScanResult result) {
        if (result.getError() != null) {
            showError(result.getError());
            return;
        }

        String status = String.format(
            "Bucket: %s\n\n" +
            "Status:\n" +
            "• Exists: %s\n" +
            "• Public Access: %s\n" +
            "• Region: %s",
            result.getBucketName(),
            result.getExists() ? "Yes" : "No",
            result.getIsPublic() ? "Yes (VULNERABLE)" : "No",
            result.getRegion()
        );

        resultText.setText(status);
        if (result.getIsPublic()) {
            showError("Warning: This bucket is publicly accessible!");
        }
    }

    private void showError(String message) {
        Snackbar.make(
            findViewById(android.R.id.content),
            message,
            Snackbar.LENGTH_LONG
        ).show();
    }
}
