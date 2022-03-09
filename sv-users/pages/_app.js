import "../styles/globals.css";
import Router from "next/router";
import NProgress from "nprogress";
import "nprogress/nprogress.css";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Head from "next/head";
import { QueryClientProvider, QueryClient, Hydrate } from "react-query";
import { useState } from "react";
// import { Provider } from "react-redux";
// import store from "redux/store";

NProgress.configure({ showSpinner: false }); //showSpinner: false
Router.events.on("routeChangeStart", () => NProgress.start());
Router.events.on("routeChangeComplete", () => NProgress.done());
Router.events.on("routeChangeError", () => NProgress.done());

function MyApp({ Component, pageProps }) {
  const [queryClient] = useState(() => new QueryClient());
  return (
    <QueryClientProvider client={queryClient}>
      <Hydrate state={pageProps.dehydratedState}>
        {/* <Provider store={store}> */}
        <Head>
          <meta name="viewport" />
          <link rel="shortcut icon" href="/favicon2.ico" />
        </Head>
        <ToastContainer theme="colored" />
        <Component {...pageProps} />
        {/* </Provider> */}
      </Hydrate>
    </QueryClientProvider>
  );
}
export default MyApp;
