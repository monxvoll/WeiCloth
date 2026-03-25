import { Routes } from '@angular/router';

export const routes: Routes = [
    {
        path: '',
        loadChildren: () =>
        import('./auth/auth-module') //Antes de cargar algo, angular ejecuta la funcion en busqueda de un paquete de codigo
            .then(m => m.AuthModule) //Arranca especificamente la clase que viene dentro d ese archivo
    },
    {
        path: 'clothes',
        loadChildren: () =>
        import('./clothes/clothes-module')
        .then(m=>m.ClothesModule)

    }
];