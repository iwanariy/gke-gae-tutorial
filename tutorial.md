# GKE / GAE tutorial

## 概要
hoge

## 事前準備
### ToDo
* Enable API

## チュートリアルの流れ
1. Google Kubernetes Engine
    1. クラスタの作成
    2. アプリケーションのデプロイ
    3. アプリケーションの確認
2. Google App Engine
    1. アプリケーションのデプロイ
    2. アプリケーションの確認

## GKE: 1. クラスタの作成
クラスタの作成  
```bash
gcloud container clusters create cluster-prod --num-nodes=1 --zone=asia-northeast1-c
```

認証情報の取得
```bash
gcloud container clusters get-credentials cluster-prod --zone asia-northeast1-c
```

## GKE: 2. アプリケーションのデプロイ

```bash
kubectl apply -f k8s_manifests/blue.yaml
```


## GKE: 3. アプリケーションの確認

```bash
kubectl get deployment
kubectl get service
```


## GAE: 1. アプリケーションのデプロイ
```bash
cd appengine
```

```bash
gcloud app deploy
```

## GAE: 2. アプリケーションの確認
`https://<project-id>.appspot.com` にアクセス。
