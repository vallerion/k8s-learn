## Part 1 - Pods

`kubectl apply -f nginx.yaml`
`kubectl port-forward nginx 3000:80`
`kubectl port-forward nginx --address 0.0.0.0 3000:80`

`kubectl logs --follow nginx`

`kubectl exec nginx -- ls`
`kubectl exec -it nginx -- bash`

`kubectl delete pod nginx`
`kubectl delete -f nginx.yaml`

## Part 2 - Deployments

`docker build -t vallerion/k8s-cource-1 .`
`docker push vallerion/k8s-cource-1`

`kubectl apply -f hellok8s.yaml`
`kubectl get deployments -o wide`
`kubectl get pods --watch`
`kubectl port-forward hellok8s-79896b8b69-pg6cg 3000:8081`
`kubectl delete pod hellok8s-79896b8b69-pg6cg`
`kubectl rollout undo deployment hellok8s`
`kubectl describe pod hellok8s-6dc69c6bf-xsvpc`

## Part 3 - Services

`kubectl get service hellok8s-svc`
`kubectl get svc`
`kubectl exec -ti hellok8s-f4447988f-r6rjm -- sh`
`kubectl describe pod nginx`
`k get pods`
`k exec -ti hellok8s-79f6946f94-tpntw -- sh`
`> env | grep HELLOK8S`

## Part 4 - Ingress

Install: `kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.6.4/deploy/static/provider/cloud/deploy.yaml`
`kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx`
`kubectl delete deployment,service --all`

## Part 5 & 6 - ConfigMap & Secrets
`k rollout restart deploy hellok8s`
`echo 'It works with a Secret' | base64`
`k exec -ti hellok8s-76d6d9d4c4-q5fjc -- env | grep SECRET`
`k get secret hellok8s-secret -o yaml`
`k delete secrets --all`
`k exec -ti hellok8s-6fc597dcb9-5z22m -- cat /secrets/SECRET_MESSAGE`

## Part 7 - Jobs
`k apply -f echo-job`
`k get jobs`
`k get pods`
`k logs echo-job-46nkd`
`kubectl delete jobs,cronjobs --all`
`k delete cronjob echo-cronjob`

## Part 8 - Namespace
`k apply -f taras-namespace.yaml`
`k apply -f nginx.yaml -n taras`
`k get pods -n taras`
`k get namespaces`
`kubectl get pods --all-namespaces`
`kubectl get all --all-namespaces`
`k -n taras exec -ti hellok8s-7c56bbd7d4-vtlsm -- sh`
`wget hellok8s-svc:8081`
`wget hellok8s-svc.default:8081`
`wget hellok8s-svc.taras:8081`

## Part 9 - Resource Usage
`kubectl describe nodes`
`k exec -ti loop-58c86f5675-t8p99 -- top`
`k describe pod loop-7b46647889-kdt6x`
`k delete pod limited-busybox --grace-period=0 --force`
`k delete limitrange memory-limit-range`

## Part 10 - Kubeconfig
`kubectl config current-context`
`k get pods --context production`

## Project
`k apply -k .`
`k get pods -o wide`
`k get nodes -o wide`
`curl 142.93.232.217:30001`
`k get svc` -> `curl 159.223.240.63` <- LoadBalancer
`kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.6.4/deploy/static/provider/cloud/deploy.yaml`
`kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission`
`kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx`
`kubectl get svc --all-namespaces -l app.kubernetes.io/name=ingress-nginx` to get external ip `134.209.136.80`
echo `134.209.136.80 nginx.do.com` > /etc/hosts
echo `134.209.136.80 hello.do.com` > /etc/hosts