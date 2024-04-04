
```shell
# 更新证书
kubeadm certs renew all
# 更新 config
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
# 查看过期时间

kubeadm certs check-expiration

```