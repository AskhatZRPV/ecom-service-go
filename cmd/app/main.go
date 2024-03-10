package main

import (
	"context"
	"ecomsvc/internal/application/account/createaccount"
	"ecomsvc/internal/application/account/deleteaccount"
	"ecomsvc/internal/application/account/updateaccount"
	"ecomsvc/internal/application/address/createaddress"
	"ecomsvc/internal/application/address/deleteaddress"
	"ecomsvc/internal/application/address/updateaddress"
	"ecomsvc/internal/application/auth/login"
	"ecomsvc/internal/application/auth/register"
	addproducttocart "ecomsvc/internal/application/cart/addproductocart"
	"ecomsvc/internal/application/cart/deletecart"
	"ecomsvc/internal/application/cart/getcart"
	"ecomsvc/internal/application/cart/updatecartitemquantity"
	"ecomsvc/internal/application/category/createcategory"
	"ecomsvc/internal/application/category/deletecategory"
	"ecomsvc/internal/application/category/updatecategory"
	"ecomsvc/internal/application/order/createorder"
	"ecomsvc/internal/application/product/createproduct"
	"ecomsvc/internal/application/product/deleteproduct"
	"ecomsvc/internal/application/product/getallproducts"
	"ecomsvc/internal/infrastructure/crosscutting/bcrypt"
	"ecomsvc/internal/infrastructure/crosscutting/pgsql_client"
	viperconfig "ecomsvc/internal/infrastructure/crosscutting/viper_config"
	"ecomsvc/internal/infrastructure/paseto_tokens"
	httprest "ecomsvc/internal/interface/http_rest"
	"ecomsvc/internal/interface/http_rest/auth/post_login"
	"ecomsvc/internal/interface/http_rest/auth/post_register"
	"ecomsvc/internal/interface/http_rest/cart/delete_deletecart"
	"ecomsvc/internal/interface/http_rest/cart/get_getsinglecart"
	"ecomsvc/internal/interface/http_rest/cart/post_additem"
	"ecomsvc/internal/interface/http_rest/cart/put_updatecart"
	"ecomsvc/internal/interface/http_rest/category/delete_deletecategory"
	"ecomsvc/internal/interface/http_rest/category/post_createcategory"
	"ecomsvc/internal/interface/http_rest/category/put_updatecategory"
	"ecomsvc/internal/interface/http_rest/common"
	"ecomsvc/internal/interface/http_rest/order/post_createorder"
	"ecomsvc/internal/interface/http_rest/product/delete_deleteproduct"
	"ecomsvc/internal/interface/http_rest/product/get_getallproducts"
	"ecomsvc/internal/interface/http_rest/product/post_createproduct"
	"ecomsvc/internal/interface/http_rest/user_account/delete_deletepayment"
	"ecomsvc/internal/interface/http_rest/user_account/post_createpayment"
	"ecomsvc/internal/interface/http_rest/user_account/put_updatebalance"
	"ecomsvc/internal/interface/http_rest/user_address/delete_deleteaddress"
	"ecomsvc/internal/interface/http_rest/user_address/post_createaddress"
	"ecomsvc/internal/interface/http_rest/user_address/put_updateaddress"

	cartitemRepo "ecomsvc/internal/infrastructure/repository/cart_item/pgsql"
	categoryRepo "ecomsvc/internal/infrastructure/repository/category/pgsql"
	inventoryRepo "ecomsvc/internal/infrastructure/repository/inventory/pgsql"
	orderRepo "ecomsvc/internal/infrastructure/repository/order/pgsql"
	orderitemRepo "ecomsvc/internal/infrastructure/repository/order_item/pgsql"
	paymentRepo "ecomsvc/internal/infrastructure/repository/payment/pgsql"
	productRepo "ecomsvc/internal/infrastructure/repository/product/pgsql"
	sessionRepo "ecomsvc/internal/infrastructure/repository/session/pgsql"
	shoppingsessionRepo "ecomsvc/internal/infrastructure/repository/shopping_session/pgsql"
	userRepo "ecomsvc/internal/infrastructure/repository/user/pgsql"
	useraccountRepo "ecomsvc/internal/infrastructure/repository/user_account/pgsql"
	useraddressRepo "ecomsvc/internal/infrastructure/repository/user_address/pgsql"

	"ecomsvc/internal/infrastructure/tx/pgsqltx"
	"os/signal"
	"syscall"
)

/*
&.................+&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&x:................&
&..;xx;+;::...........:&&&&&&&&&&&x+;:..:..::::+xxXxxx+$XXX++xxXXx:;::.:....::+&&&&&&&&&&&;.........;.:;;;+xX+..&
&&...:++;;+x;::::::.....X&&&&&:...................................................:&&&&&$....::::::::+++;;;:...&&
&&&+:::::::;;;.::::.....:&&&;.......................:...............................+&&&;.....::::..:::::;.::+&&&
&&&&&&:::;;;;;::::;+;::+::.............................:...............................:;;:::;:::::++::;;:.&&&&&&
&&&&&&x.X&&XX$Xx::;xX;..::.........::::...::::..::...............:::::....::;::........::..:Xx+::+xXxX&$X.+&&&&&&
&&&&&&&:.:+xxXXX...:.:.:............::::::::::::......:...:..:::::..::::::::::...........:::::...XXXx+;+.:&&&&&&&
&&&&&&&&$x+&$&&&.........::::.:.::::;:;;++;:;:::.:::::::::.::::::;:;:;:;:;:::::::...::;:.........X&$xX;;x&&&&&&&&
&&&&&&&&&:+$X&X:....:;;::.....::;;;xXxx;;+x$&XX$$Xx+;......:+++xX$+xxXxx++xxxXX;;::.....::;;::...:X&X$x:$&&&&&&&&
&&&&&&&&X:&&x;....::.....::x&&&&&&&&&&&&&xx+$X$$x++xx;:.:.::++;;x$x$xxxXX&&&&&&&&&&&&x;.....:::....;;X&.+&&&&&&&&
&&&&&&&&+.x$:.....:::;+&&&&&&&&&&&&&&&&&&&&+X$&$XxxxX+;:::+:xxxXX$&&XX&&&&&&&&&&&&&&&&&&&&x:........:$&.:&&&&&&&&
&&&&&&&&x:&&:......$&&&&&&&&&&&&&&&&&&&&&&&&;X$XX+++x::;:+;:;;+Xxx$X;&&&&&&&&&&&&&&&&&&&&&&&&$:.....:$&:;&&&&&&&&
&&&&&&&&$;&$:....:x&&&&&&&&&&&&&&&&&&&&&&&&&x;XXXxx+x;;::::;X++xXXX++&&&&&&&&&&&&&&&&&&&&&&&&&x:....;&&;X&&&&&&&&
&&&&&&&&$:&;...:;$&&&&&&&&&&&&&&&&&&&&&&&&&&&$&X$X++$x+::::+X++x$XX$&&&&&&&&&&&&&&&&&&&&&&&&&&&X;:...+&:X&&&&&&&&
&&&&&&&&$.:....:+&&&&&&&&&&&&&&&&&&&&&&&&&&&&$x;XXXxx;;;+;++xxX$X+x&&&&&&&&&&&&&&&&&&&&&&&&&&&&&+:....:.X&&&&&&&&
&&&&&&&&;.....::;x&&&&&&&&&&&&&&&&&&&&&&&&&&&&&XX;+;;xXx+xx+::;+X$&&&&&&&&&&&&&&&&&&&&&&&&&&&&&$+:::..:.+&&&&&&&&
&&&&&&&&..:...xxx;:+x$&&&&&&&&&&&&&&&&&&&&&&&&&;+:.::::.:::;:+:.+x&&&&&&&&&&&&&&&&&&&&&&&&&&&x;:;;;...:..&&&&&&&&
&&&&&&&&.....:xx;;:;;+&&&&&&&&&&&&&&&&&&&&&&&&&$X..:.:;::::::..:x$&&&&&&&&&&&&&&&&&&&&&&&&&;+;;::;;:..:..X&&&&&&&
&&&&&&&x.:::++:;+x;++xx$&&&&&&&&&&&&&&&&&&&&&&&xx:.....:;:::::..;+&&&&&&&&&&&&&&&&&&&&&&&Xx+x++xx;xxx::..:&&&&&&&
&&&&&&&....+$x;;;+;+;+;+x&&&&&&&&&&&&&&&&X;::.:++:.............:+;+::++$&&&&&&&&&&&&&&&&xX+Xx++++;;X$x....&&&&&&&
&&&&&&..:..+x+++++;;+&$$xxX$X$&$&&&$xx;::::::.::+;:....::::....;+:...:::::+X&&&&&&&$X$Xx+XxX+;:;;;+;;x..:..&&&&&&
&&&&&X:.::;xXx++++xXX$XXXX+x++xxxxx&XX+:;+;x+;.:+;.............:+:.:;;++++xxxXX+x+x++++xx$$X$$++x++xXX;::..&&&&&&
&&&&&:.:::;XXX;;;;X$Xx+;xxx+;;:::::::;::...:::.:+x+:.........:;+x;:::::.:::;:;:::..:;+++xx+x$$x;;;:xX$+:::.:&&&&&
&&&:..::;:X$X+++;+xXXxx++xx:........::.....:::.:++;;:.......:;+x+:::::........:......:;++;;+xXX+;+:xX$X:::....&&&
&$..;+.:;:Xxx+;;;+x++;:;+::............:.......:xxX+:........;+xx;.:::.......:........:;+;;;+++;++;+XXX;;::+;..&&
&..;+:::;+xXxXxx+;;+;::...................:.:.::++;+:.......:;;+;:::......................::::+;+++xXxx+;;::+:..&
&.;++;+x+Xx$xxxXx;:;::::.....::...........:.:;;::;++::......:+;::.::::::......:::........::::;+;+xXXx$xXxXx;;x;.&
.::;++++++xXXX$$x++;:::::..:::::::;::.......:::::;+x+::....;xx+;:::::........::;:::::::::::::+++xXXXX$Xx+xx++;+;.
.xx+++++xxXX$$Xxx;:;:;::;:;;:::::;;:......::::::::;xX+++;x:+XX;.::::.:.......::;::::::;:;+:::::;xxx$&$X+xx+++;;+.
.x+x$$$XxX$&&&$XXx+;+;++++xx+;:::;:::.::::;:::::::;+++&&&&&+++;::;::::;;;::::::::;;;;+x+;+;;+++xxX$$$X$XxXX$$++x.
:XxxXxxX$$$&&X$XXXxxx+:::::::;:::;;:::::;;::++x+;:;:;+&&&&&+::::;++x:;::::::::;;;+++;::::::;xxX$XX$$&&$$XXxxXxXX:
;$XXxxxxX$&XX$$$XXXXXXXXx+XXXx;;;:;;::::::;;;;;;+;;;+X&&&&&x+::;++;+:;::::.:::;+;++x$XXXXX$xx+xxX$&XXx$XXxxXxXxx:
:xxXXXxx$$&&&&&$$$$$&X+++x+++xXxX+x+:+;:::+;++++;+x+X&&&&&&&$x++x+++++x+;:::+++x++xXxx++++XX&$$&&&&&&&&XXXXXXXXX:
:xxxXx+xX$&&&&&&&&&&&$$XXxxxXXXXX+x+;;+xxxx++;+++xX+X&&&&&&&$x+xx+++xxx+++++++xxxXxxxxXXXx$X&&&&&&&&&&&$XXxxx$Xx:
.X$XX$$$&&&&&&&&&&&&$$&&&&&&XX$X$X&xxxxx++xx+++xxX+x&&&&&&&&&$X$x++;+xxx++xxx+XXX&&X$&&&&$&&$&&&&&&&&&&$XX$$XXXX;
:XXXX$X$$$&&&&&&&&&&&xX$$$&$$$&&&Xxx;+XXxx++++;xXxX$&&&&&&&&&$xXXx+++x;++x$x+xX$&&XXX$$$$$XX&&&&&&&&&&&$XX$xxxxx:
:$xX$x+xXX&&&&&&&&&&&$X&&&&&$X&&&&&&X+x++++:;:;+x$X&&&&&&&&&&&X$$x;::+;++++xX&&&&&&&&&&&&&X$&&&&&&&&&&&X$xXXXXXX:
:xxxxxxxxXXxX$&&&&&&&&x:;+xX&&&$X$$&&&$X+x+++;;+x+x$&&&&&&&&&&xx++;:+x;+xx$&&&&X&&&&$xxx;:+&&&&&&&&XxXXXX+xxxXXX:
:xXxx+x++xxx++XX$&&&&&&:....+xx++xxX&&&&X+;::;;:;;xx&&&&&&&&&X+;++;;:++X$&&&&$X;;++XX:...;&&&&&&$XXXxXX+x+xxxXXx:
:xxxXxxxxx+++x;;x$&&&&&&&x.....:;::;;..::;+XXX;;;;xx$$&&&&&$$Xx;++xx$$x:::..;:::;;.....X&&&&&&&X$xxx++x++xxxXXx+:
.xXX$x++xxxxxX$$Xxx&&&&&&&:...:x$$;:....;x&&$X$xx+x$&&&$X$&&&X++x++x&$&X;.....+x$+:...:&&&&&&&+x&$$xxxxx+XxxxXxx:
.xx++xXxx+;xxxx+;;;:;+X&&&&:....:;+&....:++::::X&&&$XXX&&&&XXX&&&X::;;;;:....X+;;....;&&&&Xx+;;;;;;xxxX++xxxxx$X:
:xXXxXxxx+++x;:;;:;+x+xx$&&&&...:$$&.....:;++;+&&&&&&&&&&&&&&&&&&&x;x+:::....&&&:...&&&&x+++x+;;+;++++;xXxxxxxx+.
.;x+++xx+xx++++x++++X$$$X&&&&&;...:;........++;++xXX$$&&&&&&$xxXXx;++........:....:&&&&&X$$$X+;+x+XX++xxxxXxxXx;.
:xXxx+XXxXX++xxxxx;;+:;+xXXX&&$.............;$xXx::::.x&&&x.:::::+X&+.............&&&xXxX$+x+;++x+++xx+XX+x+xxX$:
.XXXXxxx;;+xX+++xx+;x;;+++$$x&&X.............&+.........&.........;$.............+&&&$X++x:+x+;+x++xXx+xx+XXxxxx.
:X+xxxxxxxx+++xXXX::::::::$&X$&&&&...........:..........&..........;...........&&&$$$&$X::;:;:;+Xxx++++XX+++xXx+:
:+;xxxxxxxxXX&$xx+;;:;;+:;X$X&&&&&:..........+..........X..........+...........&&&&$X$Xx;;:::::++X$$XxxXxxxxX+xx:
:xxxxxxxxxx+x++:;:;;+::::xXX&&$&&&&:....................&....................:&&&$$$$X+:;+++;;:;;;+:;;;Xxxxxxx+x:
;XXX++x++xx++;+;;:::;++;;xxX$$X$&&&&....................:....................$&&&$X$$X+;+;+x+;:;;;::;x++;++x+xxX;
+XxXXXxx+++;;;+::++;:;:;;+x++Xx++&&&&$+...................................x$&&&&xxXx+;++:;;;;;;;:;:;xx+++x+xXXX$X
X&$$X$XXxxxxXxXX+++;;+;+;xx+;+xX+x&&&&&+.................................+&&&&&Xxxx;;x+x;+;;;;:+xXXX+XXXXXX$X$$&x
XX$&X$xxXxx+xxXX+;+xX$x+;;+;;+XXxxxX&&&&&;.............................;&&&&&$X+x$x++Xxx++++;++xXXXx+++x+XX&&&$$x
$&&&X&XXxxX++XXXxX++++xx+++++x++XXXXXX&&&&:............................&&&&$XXXxx++X++++++++;+++xxXxxx+$$$$$&$$$x
X&$$x$X$XXXxxXxXX$x+++;x++X+;x+xX;;xX+xX&&&&&&&&.................&&&&&&&&&XxXx++xx;+x+xXXXXX&$$$$$$Xxx+X$$$xXX$$x
X$&&$$$$xxx+xxx$X&Xx+xx+;+x++;:++;xXxXX$&&&&&&&&.................&&&&&&&&$XXxxXxxX;++++xxxxxxX$XX$XXxxx$XX$$$$$$x
+x+xxxX++XXx++;$XX+;::::::::::...::::;+xX&&&&&&&&&&XX&&&&&&$xx&&&&&&&&$Xx;+;::......::.::;;+:++xXXx;+X+xXxxxxxxxx
*/
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := viperconfig.New()
	if err != nil {
		panic(err)
	}

	tp, err := paseto_tokens.New(cfg)
	if err != nil {
		panic(err)
	}

	pgc, err := pgsql_client.New(cfg)
	if err != nil {
		panic(err)
	}

	// Repositories...
	cartItemRepo := cartitemRepo.New(pgc)
	catRepo := categoryRepo.New(pgc)
	invRepo := inventoryRepo.New(pgc)
	ordRepo := orderRepo.New(pgc)
	ordItemRepo := orderitemRepo.New(pgc)
	paymentRepo := paymentRepo.New(pgc)
	prodRepo := productRepo.New(pgc)
	sessRepo := sessionRepo.New(pgc)
	shopSessRepo := shoppingsessionRepo.New(pgc)
	userRepo := userRepo.New(pgc)
	userAccRepo := useraccountRepo.New(pgc)
	userAddRepo := useraddressRepo.New(pgc)

	// Transaction Manager...
	tx := pgsqltx.New(pgc)

	// BCrypt Hasher...
	bh := bcrypt.New()

	// UserAccount usecase...
	crAccUC := createaccount.New(userAccRepo)
	delAccUC := deleteaccount.New(userAccRepo)
	// getAccByIDUC := getaccountbyid.New(userAccRepo)
	// getAccByUIDUC := getaccountbyuserid.New(userAccRepo)
	updAccUC := updateaccount.New(userAccRepo)

	// UserAddress usecase...
	crAddUC := createaddress.New(userAddRepo)
	delAddUC := deleteaddress.New(userAddRepo)
	// getAddByIDUC := getaddressbyid.New(userAddRepo)
	// getAddByUIDUC := getaddressbyuserid.New(userAddRepo)
	updAddUC := updateaddress.New(userAddRepo)

	// Auth usecase...
	logUC := login.New(tx, bh, userRepo, tp, sessRepo)
	regUC := register.New(tx, bh, userRepo)

	// Cart usecase...
	addProdToCartUc := addproducttocart.New(tx, prodRepo, shopSessRepo, invRepo, cartItemRepo)
	// crCartUC := createcart.New(tx, shopSessRepo)
	delCartUC := deletecart.New(tx, shopSessRepo)
	getCartUC := getcart.New(tx, prodRepo, shopSessRepo, cartItemRepo)
	// remProdCartUC := removeproductfromcart.New(tx, prodRepo, shopSessRepo, cartItemRepo) //FIXME: no inventory,
	updCartItemQuanUC := updatecartitemquantity.New(tx, prodRepo, invRepo, shopSessRepo, cartItemRepo)

	// Category usecase...
	crCatUC := createcategory.New(catRepo)
	delCatUC := deletecategory.New(catRepo)
	// getCatByIDUC := getcategorybyid.New(catRepo)
	updCatUC := updatecategory.New(catRepo)

	// Order usecase...
	crOrdUC := createorder.New(tx, userAccRepo, ordRepo, ordItemRepo, paymentRepo, shopSessRepo, cartItemRepo)
	// getAllOrdUC := getallorders.New(tx, prodRepo, shopSessRepo, cartItemRepo)
	// getOrdUC := getorder.New(tx, prodRepo, shopSessRepo, cartItemRepo)
	// getusersordUC

	// Product usecase...
	crProdUC := createproduct.New(tx, prodRepo, invRepo)
	delProdUC := deleteproduct.New(tx, prodRepo, invRepo)
	getAllProdUC := getallproducts.New(tx, prodRepo, invRepo)
	// getProdByIDUC := getproductbyid.New(tx, prodRepo, invRepo)
	// updProdUC := updateproduct.New(tx, prodRepo, invRepo)

	//Handlers
	// var handlers []common.Handler
	handlers := []common.Handler{
		// Account handlers
		post_createpayment.New(crAccUC),
		delete_deletepayment.New(delAccUC),
		put_updatebalance.New(updAccUC),

		// Address handlers
		post_createaddress.New(crAddUC),
		delete_deleteaddress.New(delAddUC),
		put_updateaddress.New(updAddUC),

		// Auth
		post_login.New(logUC),
		post_register.New(regUC),

		// Cart handlers
		post_additem.New(addProdToCartUc),
		// post_createcart.New()
		delete_deletecart.New(delCartUC),
		get_getsinglecart.New(getCartUC),
		put_updatecart.New(updCartItemQuanUC),

		// Category handlers...
		post_createcategory.New(crCatUC),
		delete_deletecategory.New(delCatUC),
		// get_getcategory.New(getCatByIDUC),
		put_updatecategory.New(updCatUC),

		// Order handlers...
		post_createorder.New(crOrdUC),
		// get_getallorders.New(getAllOrdUC),
		// get_getorder.New(getOrdUC),

		post_createproduct.New(crProdUC),
		delete_deleteproduct.New(delProdUC),
		get_getallproducts.New(getAllProdUC),
		// get_getproduct.New(getProdByIDUC),
		// put_updateproduct.New(updProdUC),
	}

	httprest.New(ctx, cfg, handlers)
}
