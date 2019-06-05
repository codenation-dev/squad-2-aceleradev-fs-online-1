import React from 'react';
import { BrowserRouter, Route, Switch} from 'react-router-dom';

import Main from './pages/Main';
import UploadClients from './pages/UploadClients';

const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route path="/" exact component={Main}/>
            <Route path="/upl/:id" component={UploadClients}/>
        </Switch>
    </BrowserRouter>
);

export default Routes;