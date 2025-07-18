# Tenant operator

## k8s operator에 대해

- 네임스페이스 관리
  - 생성된 각 테넌트 개체에 대해 k8s operator는 클러스터에 네임스페이스를 생성한다.
  - 각각 테넌트에 대해 별도의 네임스페이스를 가질 수 있다.
  - 테넌트의 이메일을 각 네임스페이스에 기록한다.
- 역할 관리
  - k8s operator는 각 테넌트가 해당 네임스페이스에서 필요한 권한을 갖도록 각 네임스페이스의 Rolebinding 및 Role을 관리
- 수명 주기 관리
  - k8s operator는 테넌트 객체가 삭제될 때 Clean-up과 같은 수명 주기 이벤트를 처리
- 충돌 해결
  - 두 테넌트가 동일한 네임스페이스를 갖지 못한다.

## CRD Spec

```yaml
apiVersion: multitenancy.1eedaegon.github.io/v1
kind: Tenant
metadata:
  name: tenant-sample
spec:
  adminEmail: admin@yourdomain.com
  adminGroups:
    - tenant-sample-admins
  userGroups:
    - tenant-sample-users
    - another-group-users
  namespaces:
    - tenant-sample-ns1
    - tenant-sample-ns2
    - tenant-sample-ns3
```

## Buissiness

`spec` 섹션 아래의 `namespaces`에 언급된 대로 세 개의 네임스페이스를 생성한다.
각 네임스페이스 내에서 테넌트의 관리 그룹(adminGroups) 및 사용자 그룹(userGroups)에 대한 RoleBinding을 생성한다.
adminEmail로 각 네임스페이스에 주석을 기재.

## Requirements

- kubebuilder
- minikube

## 참고자료

- https://www.codereliant.io/hands-on-kubernetes-operator-development-part-2/
