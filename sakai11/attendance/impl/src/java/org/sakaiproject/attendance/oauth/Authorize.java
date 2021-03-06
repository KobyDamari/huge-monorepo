package org.sakaiproject.attendance.oauth;

import org.sakaiproject.attendance.export.GoogleClient;

import com.google.api.client.googleapis.auth.oauth2.GoogleAuthorizationCodeFlow;
import com.google.api.client.googleapis.auth.oauth2.GoogleTokenResponse;
import com.google.api.client.googleapis.javanet.GoogleNetHttpTransport;
import com.google.api.client.http.HttpTransport;
import com.google.api.client.json.jackson2.JacksonFactory;
import com.google.api.client.util.store.FileDataStoreFactory;
import com.google.api.services.admin.directory.DirectoryScopes;

import java.io.BufferedReader;
import java.io.File;
import java.io.InputStreamReader;


public class Authorize {
    final static String COPY_PASTE_OAUTH_TOKEN = "urn:ietf:wg:oauth:2.0:oob";

    public static void main(String[] args) {
        try {
            if (args.length != 3) {
                System.err.println("Usage: Authorize <oauth user> <oauth secret> <output name>");
                System.exit(1);
            }

            String user = args[0];
            String secret = args[1];
            String output = args[2];

            File dataStoreLocation = new File(output);
            FileDataStoreFactory store = new FileDataStoreFactory(dataStoreLocation);

            HttpTransport httpTransport = GoogleNetHttpTransport.newTrustedTransport();
            JacksonFactory jsonFactory = JacksonFactory.getDefaultInstance();

            // General idea: create the auth flow backed by a credential store;
            // check whether the store already has credentials for the user we
            // want.  If it doesn't, we go through the auth process.
            GoogleAuthorizationCodeFlow auth = new GoogleAuthorizationCodeFlow.Builder(httpTransport, jsonFactory,
                    user, secret,
                    GoogleClient.requiredScopes())
                    .setAccessType("offline")
                    .setDataStoreFactory(store)
                    .build();

            // Construct a URL for the user to hit.  This will give them a code to paste back in.
            System.err.println("Please visit the following URL:\n\n" + auth.newAuthorizationUrl().setRedirectUri(COPY_PASTE_OAUTH_TOKEN).toString() + "\n\n");

            System.err.print("Paste the code you received here: ");

            String code = new BufferedReader(new InputStreamReader(System.in)).readLine().trim();

            // Take the code we got and request a new token (access token + refresh token) from Google.
            //
            // No further copy/paste happens here, but for some reason I was
            // required to add a redirect_uri to the request, even though
            // it's not obviously used...
            GoogleTokenResponse response = auth.newTokenRequest(code).setRedirectUri(COPY_PASTE_OAUTH_TOKEN).execute();

            // Store the token in our credential store for laster
            auth.createAndStoreCredential(response, user);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
