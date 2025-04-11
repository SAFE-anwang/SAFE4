// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main SAFE network.
var MainnetBootnodes = []string{
	"enode://c1c0779cb05f85cccd6ca26225df87b36b6501d72b0687d3ca1cd0ffec05dbfb3d1b44264103a7a72b4c1d2291a3869dc0318edfed8a3c3089a3a8e8024d0799@39.104.90.76:30303",
	"enode://01001203dbc45d6078701756e20749fb5eea018c0ec68871cb2090416492a6bb8d55aa73307965adee9fb9024d9c9cb91615c7ff241eb87ddfb5007919141f6d@120.78.227.96:30303",
	"enode://0a1f0e1c5a166f9afe32333670a57a98122f8d90a81057a135be638d75e69470dd312517e161fa73081269724679bb5730ade2034693ff01be5858f41b1b5b3a@39.104.200.133:30303",
	"enode://7e6f55ba0659e6a68a4384e8097731fa6f204409b8b63319c2da2a3e04f2b03c5268fa36ce297d1bdf12d907273728536e59d17bc77668047ad62b3083815f4d@106.14.66.206:30303",
	"enode://bb3ca0a7cf94f03e5be283a0fb3b78891ba1bb3cd40350055ac5a43d9cc706c08f28d4a57ca2dc4b598091228492b17f2f488f50add947467b5fa038b74ec1b3@114.215.31.37:30303",
	"enode://c5db0aa4bb9d37d20c8967f4d75e1b8d60d27c4f2a4666817457576786cb16d2d398f35add358e08f51d908f712f8aa480351802084ea800289d90bb718abf4c@47.96.254.235:30303",
	"enode://002d953f34f90f289ea88cc8f0bcf887eb0dfb5cb94b74a6db284efca90dc3840dd3ddab086aa1f1d5bf8e2ef2f0e7d43284b4b3f8c80377a78ada4144d9388f@192.53.112.232:30303",
	"enode://2aa0b09ba5e87a1550dddb134bc398e7d2ef3835f541f3159ae8d4ec1ddf3f380785bb6be4255cb3dc86c9fdd67e24ca110aab6e74711d759cda2c4d076c6f8b@139.162.10.148:30303",
	"enode://1a6b6058517485c739c96a4a60e56c70b37330808158884e6c8693dc04816cff3dff350b7a2fe25b7b492be20969336631665d078b0d91736cca330b60bd51ee@172.105.209.98:30303",
	"enode://8f0c900dd3b8fe0d526a712b05e4b76e3d771e450fd82f6af3e7926ed00e468966130f957f9c54bd7f3daf2722e99692c6cbe05cd260c506bd2481b3ad9d6449@172.105.235.94:30303",
	"enode://b00fe76dee72630eb0afd4f1d5a5faffeb6a9070bc334c9eba289d6c4aa27c6d30e1deccf684763b28d0da6908f2228d061647b40198a0b1f6a133b60ae8d671@172.104.28.167:30303",
	"enode://9cb23565d5a0b58fce185d2d86b9afbd583dbf47416f959bbc257d66cdd99376acd5d28b52cc9e77b87054685fa62e9ef39832d0d15b63a0bdb2a8df553cc45e@139.162.196.118:30303",
	"enode://c78817e4b636a8473887754d656a1ce49a92ad5d825e8d454360d475ab166e7e678cfbfb2286ec36c636495b0ffd3f6daf80787ee7bc59955eab45242f5c80d2@212.111.40.32:30303",
	"enode://5090ab3b35d5303a2f781e0c121e76e3544b4b1f27dff3f0cbc141ef39e0357c4d8cf5d5c7eecfec675589c6d7a9bab3519a2fb07122f5f9f2456787dd2674ef@139.162.142.45:30303",
	"enode://05d0801d0087ce844391dda038ceee0d70cd620790db9e59ab6c36abf6f49da802ca2d9a8d8f6258ae2ed2e6846c9296897ef36944fc434594a9ae7665ff69d8@192.46.232.91:30303",
}

// the SAFE test network.
var SafeTestBootnodes = []string{
	// SAFE test chain Go Bootnodes
	"enode://57e52467b3bdfbe62b1c47a16be714ca7e5ae117f62453142637019ed2b8daa707860a004d381ad3f73f815f61bdbc688f2f40143239e324a8666d075c95f3f0@172.104.85.174:30303", // sn1
	"enode://1ac4e9530ec9d6fb1b561e3b6af91768db767be6ccb236ee545944ae1cc578695d264d143216b2ad5447bb875c92b7d247034c74d98ce85923b985d1365ba572@139.162.103.9:30303",  // sn2
	"enode://27dcaa3373400bdaf41bbc297f1e1db82ca4c4a5501731563c202be5c5b5b605a166ff820daecd29f7dde161677ff0fd9ac8e476602bfea00e0cbf81ad0e2a87@172.105.112.125:30303",// sn3
	"enode://f4c04849771c2f0c7e5edfb293dda9b68b46076d9d4860cc1204ee498c67be0699a0b8ee55c146baf97da2393a2c9d8cf72f6c57d091ad62af01ffd0478b78cd@47.52.9.168:30303",    // sn4
	"enode://93edcb1c25dd16f8f440c12401c7bbfe26e7577e22289ca2ac7a1c8857018b6448ce7e4db799f943bb756d50f965f0af7b8a7690b425c780048a88cf31ce18c9@172.105.201.79:30303", // sn5
	"enode://14468e09fd9a61bd09104669f0725c25f4829ae1835c65867a8799f4af92dcaf3b333b2757cc0ced7c00c89d7c3c6f5d3c0bc6a4d1674775bf10324758ddc939@47.75.17.223:30303",   // sn6
	"enode://20c8af2dd90a8f5585cf4ec2676f85a309e3f7cef399dfb1c5b01763c822f5881e5aab98091709b8331249d564e15854bfc152c14725c766e547d55d375714a0@47.88.247.232:30303",  // sn7
	"enode://2ffaec43add0159d6e7e52f1c863f02762054caa64fcc9f05e8364e2c9e99ab8861575e9880ff01a4d98f9b4626d46133f3a43205093e4ac0c816d0112f6e32f@47.89.208.160:30303",  // sn8
	"enode://58244818dfbcdbc89905d315bb568eee4e96bc3f94a628fcee1558c70d29de5f0b98290cb71a4a8c63276e8cbe5d3385ff5f7e5baaf9e44215e755ab2f7132d8@172.105.6.192:30303",  // sn9
	"enode://d39fc9ed12000b2ea3b5463936958702f20a939405aae28e39463c8b66d78bb07baf7fe59370f6037849f2bd363b1bee3301d6b0bf8349abfbcafb5fccdceab2@172.105.24.28:30303",  // sn10
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

var V5Bootnodes = []string{
	// Teku team's bootnode
	"enr:-KG4QOtcP9X1FbIMOe17QNMKqDxCpm14jcX5tiOE4_TyMrFqbmhPZHK_ZPG2Gxb1GE2xdtodOfx9-cgvNtxnRyHEmC0ghGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQDE8KdiXNlY3AyNTZrMaEDhpehBDbZjM_L9ek699Y7vhUJ-eAdMyQW_Fil522Y0fODdGNwgiMog3VkcIIjKA",
	"enr:-KG4QDyytgmE4f7AnvW-ZaUOIi9i79qX4JwjRAiXBZCU65wOfBu-3Nb5I7b_Rmg3KCOcZM_C3y5pg7EBU5XGrcLTduQEhGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQ2_DUbiXNlY3AyNTZrMaEDKnz_-ps3UUOfHWVYaskI5kWYO_vtYMGYCQRAR3gHDouDdGNwgiMog3VkcIIjKA",
	// Prylab team's bootnodes
	"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg",
	"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA",
	"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg",
	// Lighthouse team's bootnodes
	"enr:-IS4QLkKqDMy_ExrpOEWa59NiClemOnor-krjp4qoeZwIw2QduPC-q7Kz4u1IOWf3DDbdxqQIgC4fejavBOuUPy-HE4BgmlkgnY0gmlwhCLzAHqJc2VjcDI1NmsxoQLQSJfEAHZApkm5edTCZ_4qps_1k_ub2CxHFxi-gr2JMIN1ZHCCIyg",
	"enr:-IS4QDAyibHCzYZmIYZCjXwU9BqpotWmv2BsFlIq1V31BwDDMJPFEbox1ijT5c2Ou3kvieOKejxuaCqIcjxBjJ_3j_cBgmlkgnY0gmlwhAMaHiCJc2VjcDI1NmsxoQJIdpj_foZ02MXz4It8xKD7yUHTBx7lVFn3oeRP21KRV4N1ZHCCIyg",
	// EF bootnodes
	"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg",
	"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg",
	"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg",
	"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
