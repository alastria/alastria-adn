DISTRIBUTED CERTIFICATIONS

According to our sdk confs we need following certifications for each entity to be distributed in seperated location/server. With this setting, we can assure that each entity will only use its own private keys and some of public keys of other parties(depends on its role in our application) 


```
Entity orgAdm
entityCoreAdm_crypto_configs/
├── ordererOrganizations
│   └── alastria.com
│       ├── tlsca
│       │   ├── cbc0305ef5144eff2010f4138b4dfba23b3d2a26386b3e7a76cffd767f3adefd_sk
│       │   └── tlsca.alastria.com-cert.pem
│       └── users
│           └── Admin@alastria.com
│               └── msp
│                   ├── admincerts
│                   │   └── Admin@alastria.com-cert.pem
│                   ├── cacerts
│                   │   └── ca.alastria.com-cert.pem
│                   ├── keystore
│                   │   └── 90464cbaf0147b65c07b92a06d95f66c447421ee6ca925d14f0c3a25eeaddd01_sk
│                   ├── signcerts
│                   │   └── Admin@alastria.com-cert.pem
│                   └── tlscacerts
│                       └── tlsca.alastria.com-cert.pem
└── peerOrganizations
    ├── coreAdm.alastria.com
    │   ├── tlsca
    │   │   ├── 75dc9985ec0c35f33de25301bd529e12c88c0619d643b58cf7ab3b6d5d1dee5e_sk
    │   │   └── tlsca.coreAdm.alastria.com-cert.pem
    │   └── users
    │       └── Admin@coreAdm.alastria.com
    │           ├── msp
    │           │   ├── admincerts
    │           │   │   └── Admin@coreAdm.alastria.com-cert.pem
    │           │   ├── cacerts
    │           │   │   └── ca.coreAdm.alastria.com-cert.pem
    │           │   ├── keystore
    │           │   │   └── d06595b8155da432c86989b873d67bd470c3dfc1655b95f8ed7b4f9d70b50db3_sk
    │           │   ├── signcerts
    │           │   │   └── Admin@coreAdm.alastria.com-cert.pem
    │           │   └── tlscacerts
    │           │       └── tlsca.coreAdm.alastria.com-cert.pem
    │           └── tls
    │               ├── ca.crt
    │               ├── server.crt
    │               └── server.key
    ├── org1.alastria.com
    │   ├── tlsca
    │   │   ├── e1a9cf0ac6f3075ea0d02eb755a1f04da160cd59ad1bd9f191cd3b8716c35fc4_sk
    │   │   └── tlsca.org1.alastria.com-cert.pem
    │   └── users
    │       └── Admin@org1.alastria.com
    │           └── msp
    │               ├── admincerts
    │               │   └── Admin@org1.alastria.com-cert.pem
    │               ├── cacerts
    │               │   └── ca.org1.alastria.com-cert.pem
    │               ├── keystore
    │               │   └── 85ae2fc501859bea4ac9c358d60159ae01d468427da74b8ccb3af6df6f2e0b83_sk
    │               ├── signcerts
    │               │   └── Admin@org1.alastria.com-cert.pem
    │               └── tlscacerts
    │                   └── tlsca.org1.alastria.com-cert.pem
    └── org2.alastria.com
        ├── tlsca
        │   ├── b2937e0721fb6109de72c880496b922a6d6341b4c71762fa77e33c3da7b8075b_sk
        │   └── tlsca.org2.alastria.com-cert.pem
        └── users
            └── Admin@org2.alastria.com
                └── msp
                    ├── admincerts
                    │   └── Admin@org2.alastria.com-cert.pem
                    ├── cacerts
                    │   └── ca.org2.alastria.com-cert.pem
                    ├── keystore
                    │   └── 457544e165a967ad7af7f102c17496221bd9d29b4e2b458dcbe917428caed2ac_sk
                    ├── signcerts
                    │   └── Admin@org2.alastria.com-cert.pem
                    └── tlscacerts
                        └── tlsca.org2.alastria.com-cert.pem



Entity org1
entity1_crypto_configs/
├── ordererOrganizations
│   └── alastria.com
│       ├── tlsca
│       │   ├── cbc0305ef5144eff2010f4138b4dfba23b3d2a26386b3e7a76cffd767f3adefd_sk
│       │   └── tlsca.alastria.com-cert.pem
│       └── users
│           └── Admin@alastria.com
│               └── msp
│                   ├── admincerts
│                   │   └── Admin@alastria.com-cert.pem
│                   ├── cacerts
│                   │   └── ca.alastria.com-cert.pem
│                   ├── keystore
│                   │   └── 90464cbaf0147b65c07b92a06d95f66c447421ee6ca925d14f0c3a25eeaddd01_sk
│                   ├── signcerts
│                   │   └── Admin@alastria.com-cert.pem
│                   └── tlscacerts
│                       └── tlsca.alastria.com-cert.pem
└── peerOrganizations
    └── org1.alastria.com
        ├── tlsca
        │   ├── e1a9cf0ac6f3075ea0d02eb755a1f04da160cd59ad1bd9f191cd3b8716c35fc4_sk
        │   └── tlsca.org1.alastria.com-cert.pem
        └── users
            └── Admin@org1.alastria.com
                ├── msp
                │   ├── admincerts
                │   │   └── Admin@org1.alastria.com-cert.pem
                │   ├── cacerts
                │   │   └── ca.org1.alastria.com-cert.pem
                │   ├── keystore
                │   │   └── 85ae2fc501859bea4ac9c358d60159ae01d468427da74b8ccb3af6df6f2e0b83_sk
                │   ├── signcerts
                │   │   └── Admin@org1.alastria.com-cert.pem
                │   └── tlscacerts
                │       └── tlsca.org1.alastria.com-cert.pem
                └── tls
                    ├── ca.crt
                    ├── server.crt
                    └── server.key
                    
Entity org2
entity2_crypto_configs/
├── ordererOrganizations
│   └── alastria.com
│       ├── tlsca
│       │   ├── cbc0305ef5144eff2010f4138b4dfba23b3d2a26386b3e7a76cffd767f3adefd_sk
│       │   └── tlsca.alastria.com-cert.pem
│       └── users
│           └── Admin@alastria.com
│               └── msp
│                   ├── admincerts
│                   │   └── Admin@alastria.com-cert.pem
│                   ├── cacerts
│                   │   └── ca.alastria.com-cert.pem
│                   ├── keystore
│                   │   └── 90464cbaf0147b65c07b92a06d95f66c447421ee6ca925d14f0c3a25eeaddd01_sk
│                   ├── signcerts
│                   │   └── Admin@alastria.com-cert.pem
│                   └── tlscacerts
│                       └── tlsca.alastria.com-cert.pem
└── peerOrganizations
    └── org2.alastria.com
        ├── tlsca
        │   ├── b2937e0721fb6109de72c880496b922a6d6341b4c71762fa77e33c3da7b8075b_sk
        │   └── tlsca.org2.alastria.com-cert.pem
        └── users
            └── Admin@org2.alastria.com
                ├── msp
                │   ├── admincerts
                │   │   └── Admin@org2.alastria.com-cert.pem
                │   ├── cacerts
                │   │   └── ca.org2.alastria.com-cert.pem
                │   ├── keystore
                │   │   └── 457544e165a967ad7af7f102c17496221bd9d29b4e2b458dcbe917428caed2ac_sk
                │   ├── signcerts
                │   │   └── Admin@org2.alastria.com-cert.pem
                │   └── tlscacerts
                │       └── tlsca.org2.alastria.com-cert.pem
                └── tls
                    ├── ca.crt
                    ├── server.crt
                    └── server.key
                    
```
