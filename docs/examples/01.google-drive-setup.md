# Setting Up Google Drive
## Detailed Guide to Google Drive Integration

---

### 1. Registering the Application
* Go to the [Google Cloud](https://cloud.google.com/);
* Access the console;
* Open the menu and navigate to the **APIs & Services tab**;
* Create a new project;
* Set the newly created project as the current one;
* Go to the **OAuth consent screen tab**;
* Choose `External` and create an application.

### 2. Creating a Service Account and Its Secret Key
* Go to the **Credentials** tab;
* Click `CREATE CREDENTIALS`;
* Select **Service account** and create an account;
* After successful creation, the created account will be displayed on the Credentials tab.
Next, you need to create a secret key for the account (click on edit, go to the **KEYS** -> `ADD KEY` -> `CREATE NEW KEY` -> `JSON` -> `CREATE`);
* After completing all the aforementioned steps, a .json file will be downloaded. You need to move this file to the `secrets` directory.
* Set the value of the environment variable `GDRIVE_CREDENTIALS` in the .env file.

### 3. Connecting to the Google Drive API
* While on the **APIs & Services** tab, navigate to the Library tab;
* Search for **"Google Drive API"**;
* The Google Drive API should appear in the search results. Click on the link and press `ENABLE`;

---
   
#### After completing all the steps described above, the service will provide the ability to obtain links to CSV files with the history of operations.