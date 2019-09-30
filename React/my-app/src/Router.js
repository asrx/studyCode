import React from 'react';
import {HashRouter, Route, Switch} from 'react-router-dom';
import Home from './views/Home';
import Foo from './views/Foo';
import Bar from './views/Bar';


// const BasicRoute = () => (
//     <HashRouter>
//         <Switch>
//             <Route exact path="/" component={Home}/>
//             <Route exact path="/foo" component={Foo}/>
//             <Route exact path="/bar" component={Bar}/>
//         </Switch>
//     </HashRouter>
// );


export default {
    Foo,
    Bar,
    Home
};
