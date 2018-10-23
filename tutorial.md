# GKE / GAE tutorial
## チュートリアルの流れ
1. Google Kubernetes Engine
    1. クラスタの作成
    2. アプリケーションのデプロイ、確認
2. Google App Engine
    1. アプリケーションのデプロイ、確認


## 事前準備
### ProjectID の指定
```bash
gcloud config set project PROJECT_ID
```

### APIの有効化  
```bash
gcloud services enable appengine.googleapis.com container.googleapis.com
```


## GKE: 1. クラスタの作成
### クラスタの作成  
```bash
gcloud container clusters create cluster-onprem --num-nodes=1 --zone=asia-northeast1-c --async
```

```bash
gcloud container clusters create cluster-cloud  --num-nodes=1 --zone=asia-northeast1-c --async
```

[Kubernetes Engine > Clusters](https://console.cloud.google.com/kubernetes/list) で結果を確認してください。 


## GKE: 2. アプリケーションのデプロイ
### 認証情報の取得  
```bash
gcloud container clusters get-credentials cluster-onprem --zone asia-northeast1-c
```

### デプロイ  
```bash
cd ~/gke-gae-tutorial
```

```bash
kubectl apply -f k8s_manifests/userservice.yaml
```

### 確認
```bash
kubectl get service
```
http://[External-IP]/ にアクセスして、アプリケーションを確認。  


上記を、 **cluster-cloud** についても実施してください。  


## GAE: 1. アプリケーションのデプロイ、確認
### デプロイ
```bash
cd ~/gke-gae-tutorial/appengine
```

```bash
gcloud app deploy
```

**Tips**: 初回利用時は、region選択が表示されるため、適宜選択してください。  

### 確認
[URL](https://<project-id>.appspot.com) にアクセス。


## Clean up
### GKE  
```bash
gcloud container clusters delete cluster-onprem cluster-cloud --zone asia-northeast1-c
```

### GAE   
```
```
