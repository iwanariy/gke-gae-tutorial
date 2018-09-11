# GKE / GAE tutorial

## 概要
hoge

## 事前準備
### ToDo
* Enable API

## チュートリアルの流れ
1. Google Kubernetes Engine
    1. クラスタの作成
    2. アプリケーションのデプロイ、確認
2. Google App Engine
    1. アプリケーションのデプロイ、確認

## GKE: 1. クラスタの作成
ProjectID の指定
```bash
gcloud config set project PROJECT_ID
```

クラスタの作成  
```bash
gcloud container clusters create cluster-onprem --num-nodes=1 --zone=asia-northeast1-c
```

認証情報の取得
```bash
gcloud container clusters get-credentials cluster-onprem --zone asia-northeast1-c
```

## GKE: 2. アプリケーションのデプロイ
### デプロイ
```bash
kubectl apply -f k8s_manifests/blue.yaml
```

### 確認
```bash
kubectl get deployment
```

```bash
kubectl get service
```
http://[External-IP]/ にアクセスして、アプリケーションを確認。  

## GAE: 1. アプリケーションのデプロイ、確認
### デプロイ
```bash
cd appengine
```

```bash
gcloud app deploy
```

### 確認
`https://<project-id>.appspot.com` にアクセス。


## Congratulations

<walkthrough-conclusion-trophy></walkthrough-conclusion-trophy>