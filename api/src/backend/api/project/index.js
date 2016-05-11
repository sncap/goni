import passport from 'backend/core/passport';
import {
  getProject,
  getProjectList,
} from './core';
import {
  Router,
} from 'express';
const router = Router();

// Project
router
  .route('/project/:id')
  .get(
    passport.authenticate('bearer'),
    async(req, res) => {
      try {
        const project = await getProject(req.params.id, req.user.id);
        if (project) {
          res.send(project);
        } else {
          res.sendStatus(404);
        }
      } catch (error) {
        res.sendStatus(500);
      }
    });

// Projects list
router
  .route('/projects')
  .get(
    passport.authenticate('bearer'),
    async(req, res) => {
      try {
        const projects = await getProjectList(req.user.id);
        res.send(projects);
      } catch (error) {
        res.sendStatus(500);
      }
    });

// Goni server list
router
  .route('/projects/goni/:id/server/list')
  .get(
    passport.authenticate('bearer'),
    async(req, res) => {

    });

// Goni server detail (Application View)
router
  .route('/projects/goni/:id/server/:serverid')
  .get(
    passport.authenticate('bearer'),
    async(req, res) => {

    });

export default router;