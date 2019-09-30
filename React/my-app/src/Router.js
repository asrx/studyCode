import React from 'react';
import {HashRouter, Route, Switch} from 'react-router-dom';
import Home from './views/Home';
import Foo from './views/Foo';
import Bar from './views/Bar';


const BasicRoute = () => (
    <div>
        <Route exact path="/" component={Home}/>
        <Route path="/foo" component={Foo}/>
        <Route path="/bar" component={Bar}/>
    </div>
);


export default BasicRoute;
