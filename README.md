# MintCocoa GitOps Demo App

Small HTTP application used to demonstrate the MintCocoa dev-first GitOps promotion flow.

Pipeline smoke test marker: 2026-06-13.

Flow:

```text
push to this app repo
  -> GitHub Actions builds ghcr.io/mint-cocoa/mintcocoa-gitops-demo-app:<commit-sha>
  -> GitHub Actions opens a staging overlay PR in mintcocoa-ops
  -> Argo CD syncs staging/home k3s after merge
  -> mintcocoa-ops promotion workflow copies the verified tag to prod
  -> Argo CD syncs OKE production after merge
```

Endpoints:

- `/healthz`
- `/`
